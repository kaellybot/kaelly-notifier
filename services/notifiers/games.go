package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
)

func (service *Impl) gameNews(ctx amqp.Context, message *amqp.RabbitMQMessage,
	game constants.AnkamaGame) {
	content := mappers.MapGameNews(message.GetNewsGameMessage(), game)
	// The reporting channel is shared, but each report is sent by its game's bot.
	service.discordService.
		SendMessage(ctx.CorrelationID, message.GetGame(), service.reportingChannelID, content)
}
