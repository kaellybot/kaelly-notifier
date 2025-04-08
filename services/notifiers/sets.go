package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
)

func (service *Impl) setNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	content := mappers.MapSetNews(message.NewsSetMessage, message.Game)
	service.discordService.
		SendMessage(ctx.CorrelationID, service.reportingChannelID, content)
}
