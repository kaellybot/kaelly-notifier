package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) feedNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	service.lock.Lock()
	defer service.lock.Unlock()

	feedTypeID := message.NewsRSSMessage.Type
	feedWebhooks, errGet := service.webhookRepo.
		GetFeedWebhooks(feedTypeID, message.Game, message.Language)
	if errGet != nil {
		log.Error().Err(errGet).
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogEntityID, feedTypeID).
			Str(constants.LogGame, message.Game.String()).
			Str(constants.LogLocale, message.Language.String()).
			Msg("Cannot retrieve feed webhooks, ignoring the feed occurence")
		return
	}

	content := mappers.MapFeed(message.NewsRSSMessage, message.Language)
	webhooks := make([]*constants.Webhook, 0)
	for _, feedWebhook := range feedWebhooks {
		webhooks = append(webhooks, &feedWebhook.Webhook)
	}

	dispatched := service.dispatch(content, &entities.WebhookFeed{}, webhooks)
	log.Info().
		Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogEntityID, feedTypeID).
		Str(constants.LogGame, message.Game.String()).
		Str(constants.LogLocale, message.Language.String()).
		Int(constants.LogWebhookCount, len(feedWebhooks)).
		Int(constants.LogSucceededWebhookCount, dispatched).
		Msg("Feed published!")
}
