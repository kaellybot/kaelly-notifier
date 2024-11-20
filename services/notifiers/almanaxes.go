package notifiers

import (
	"sync"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) almanaxNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	service.lock.Lock()
	defer service.lock.Unlock()

	almanaxes := message.NewsAlmanaxMessage.Almanaxes

	var wg sync.WaitGroup
	wg.Add(len(almanaxes))
	for _, almanax := range almanaxes {
		go func() {
			defer wg.Done()
			service.dispatchAlmanax(ctx, almanax, message.Game)
		}()
	}

	wg.Wait()
}

func (service *Impl) dispatchAlmanax(ctx amqp.Context,
	almanax *amqp.NewsAlmanaxMessage_I18NAlmanax, game amqp.Game) {
	almanaxWebhooks, errGet := service.webhookRepo.
		GetAlmanaxWebhooks(game, almanax.Locale)
	if errGet != nil {
		log.Error().Err(errGet).
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogGame, game.String()).
			Str(constants.LogLocale, almanax.Locale.String()).
			Msg("Cannot retrieve almanax webhooks, ignoring the occurence")
		return
	}

	content := mappers.MapAlmanax(almanax)
	webhooks := make([]*constants.Webhook, 0)
	for _, almanaxWebhook := range almanaxWebhooks {
		webhooks = append(webhooks, &almanaxWebhook.Webhook)
	}

	dispatched := service.dispatch(content, webhooks)
	log.Info().
		Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogGame, game.String()).
		Str(constants.LogLocale, almanax.Locale.String()).
		Int(constants.LogWebhookCount, len(almanaxWebhooks)).
		Int(constants.LogSucceededWebhookCount, dispatched).
		Msg("Almanax published!")
}
