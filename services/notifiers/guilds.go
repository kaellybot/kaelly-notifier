package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) guildNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	var content string
	newsGuild := message.NewsGuildMessage
	switch newsGuild.Event {
	case amqp.NewsGuildMessage_CREATE:
		content = mappers.MapGuildCreateNews(newsGuild)
	case amqp.NewsGuildMessage_DELETE:
		content = mappers.MapGuildDeleteNews(newsGuild)
	case amqp.NewsGuildMessage_UNKNOWN:
		fallthrough
	default:
		log.Warn().
			Str(constants.LogEvent, newsGuild.Event.String()).
			Msg("Guild event not handled, ignoring it")
		return
	}

	service.discordService.SendMessage(ctx.CorrelationID,
		service.reportingChannelID, content)
}
