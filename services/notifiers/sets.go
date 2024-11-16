package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) setNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	content := mappers.MapSetNews(message.NewsSetMessage, message.Game, message.Language)
	errPub := service.discordService.
		PublishWebhook(service.internalWebhookID, service.internalWebhookToken, content)
	if errPub != nil {
		log.Error().Err(errPub).
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msg("Cannot publish internal set news, ignoring it")
	}
}
