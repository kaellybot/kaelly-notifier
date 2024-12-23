package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) twitterNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	twitterAccountID := message.NewsTwitterMessage.TwitterId
	twitterWebhooks, errGet := service.webhookRepo.GetTwitterWebhooks(twitterAccountID)
	if errGet != nil {
		log.Error().Err(errGet).
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogEntityID, twitterAccountID).
			Str(constants.LogGame, message.Game.String()).
			Str(constants.LogLocale, message.Language.String()).
			Msg("Cannot retrieve twitter webhooks, ignoring the tweet occurence")
		return
	}

	content := mappers.MapTweet(message.NewsTwitterMessage, message.Language)
	webhooks := make([]*constants.Webhook, 0)
	for _, twitterWebhook := range twitterWebhooks {
		webhooks = append(webhooks, &twitterWebhook.Webhook)
	}

	service.dispatch(ctx.CorrelationID, content, webhooks)
	log.Info().
		Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogEntityID, twitterAccountID).
		Str(constants.LogGame, message.Game.String()).
		Str(constants.LogLocale, message.Language.String()).
		Int(constants.LogWebhookCount, len(twitterWebhooks)).
		Msg("Tweet published!")
}
