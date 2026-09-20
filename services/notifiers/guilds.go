package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) guildNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	var content string
	newsGuild := message.GetNewsGuildMessage()
	switch newsGuild.GetEvent() {
	case amqp.NewsGuildMessage_CREATE:
		content = mappers.MapGuildCreateNews(newsGuild)
	case amqp.NewsGuildMessage_DELETE:
		content = mappers.MapGuildDeleteNews(newsGuild)
	case amqp.NewsGuildMessage_UNKNOWN:
		fallthrough
	default:
		log.Warn().
			Str(constants.LogEvent, newsGuild.GetEvent().String()).
			Str(constants.LogGame, message.GetGame().String()).
			Msg("Guild event not handled, ignoring it")
		return
	}

	// The reporting channel is shared, but each report is sent by its game's bot, so
	// the message author identifies the game.
	service.discordService.SendMessage(ctx.CorrelationID, message.GetGame(),
		service.reportingChannelID, content)
}
