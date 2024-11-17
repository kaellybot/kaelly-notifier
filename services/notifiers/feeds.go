package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) feedNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	feedTypeID := message.NewsRSSMessage.Type
	date := message.NewsRSSMessage.Date.AsTime()
	webhooks, err := service.feedRepo.
		Get(feedTypeID, message.Game, message.Language, date)
	if err != nil {
		log.Error().Err(err).
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msg("Cannot retrieve feed webhooks, ignoring the feed occurence")
		return
	}

	content := mappers.MapFeed(message.NewsRSSMessage)
	failedWebhooks := make([]*entities.WebhookFeed, 0)
	for _, webhook := range webhooks {
		errPub := service.discordService.
			PublishWebhook(webhook.WebhookID, webhook.WebhookToken, content)
		if errPub != nil {
			failedWebhooks = append(failedWebhooks, webhook)
		}
	}

	log.Info().
		Int(constants.LogWebhookCount, len(webhooks)).
		Int(constants.LogSucceededWebhookCount, len(webhooks)-len(failedWebhooks)).
		Msg("Feed published!")

	// TODO treat failed webhooks
}
