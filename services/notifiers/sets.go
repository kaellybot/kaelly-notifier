package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
)

func (service *Impl) setNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	content := mappers.MapSetNews(message.NewsSetMessage, message.Game, message.Language)
	service.discordService.PublishWebhook(ctx.CorrelationID, service.internalWebhookID,
		service.internalWebhookToken, content)
}
