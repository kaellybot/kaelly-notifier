package notifiers

import (
	"sync"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) almanaxNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	almanaxes := withFallback(message.NewsAlmanaxMessage.Almanaxes)

	var wg sync.WaitGroup
	wg.Add(len(almanaxes))
	for _, almanax := range almanaxes {
		go func() {
			defer wg.Done()
			service.dispatchAlmanax(ctx, almanax, message.NewsAlmanaxMessage.Source, message.Game)
		}()
	}

	wg.Wait()
}

func (service *Impl) dispatchAlmanax(ctx amqp.Context,
	almanax *amqp.NewsAlmanaxMessage_I18NAlmanax, source *amqp.Source, game amqp.Game) {
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

	content := mappers.MapAlmanax(almanax, source, service.emojiService)
	webhooks := make([]*constants.Webhook, 0)
	for _, almanaxWebhook := range almanaxWebhooks {
		webhooks = append(webhooks, &almanaxWebhook.Webhook)
	}

	service.dispatch(ctx.CorrelationID, content, webhooks)
	log.Info().
		Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogGame, game.String()).
		Str(constants.LogLocale, almanax.Locale.String()).
		Int(constants.LogWebhookCount, len(almanaxWebhooks)).
		Msg("Almanax published!")
}

// Ensure missing Almanax for specific language are always set with fallback language.
// If fallback is missing too, language is ignored.
func withFallback(almanaxes []*amqp.NewsAlmanaxMessage_I18NAlmanax) []*amqp.NewsAlmanaxMessage_I18NAlmanax {
	var fallback *amqp.Almanax
	for _, almanax := range almanaxes {
		if almanax.Almanax != nil && almanax.Locale == constants.DefaultAMQPLocale {
			fallback = almanax.Almanax
			break
		}
	}

	result := make([]*amqp.NewsAlmanaxMessage_I18NAlmanax, 0)
	for _, almanax := range almanaxes {
		if almanax.Almanax != nil {
			result = append(result, &amqp.NewsAlmanaxMessage_I18NAlmanax{
				Almanax: almanax.Almanax,
				Locale:  almanax.Locale,
			})
		} else {
			if fallback != nil {
				log.Warn().
					Str(constants.LogLocale, almanax.Locale.String()).
					Msg("Almanax not available in this language, switching on fallback")

				result = append(result, &amqp.NewsAlmanaxMessage_I18NAlmanax{
					Almanax: fallback,
					Locale:  almanax.Locale,
				})
			} else {
				log.Error().
					Str(constants.LogLocale, almanax.Locale.String()).
					Msg("Cannot deliver almanax, fallback was not set")
			}
		}
	}

	return result
}
