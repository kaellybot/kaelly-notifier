package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
)

func (service *Impl) gameNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	content := mappers.MapGameNews(message.NewsGameMessage, message.Game)
	service.discordService.PublishWebhook(ctx.CorrelationID, service.internalWebhookID,
		service.internalWebhookToken, content)
}
