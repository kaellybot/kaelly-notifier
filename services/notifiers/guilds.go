package notifiers

import (
	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) guildNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	var content *discordgo.WebhookParams
	newsGuild := message.NewsGuildMessage
	switch newsGuild.Event {
	case amqp.NewsGuildMessage_CREATE:
		content = mappers.MapGuildCreateNews(newsGuild, message.Language)
	case amqp.NewsGuildMessage_DELETE:
		content = mappers.MapGuildDeleteNews(newsGuild, message.Language)
	case amqp.NewsGuildMessage_UNKNOWN:
		fallthrough
	default:
		log.Warn().
			Str(constants.LogEvent, newsGuild.Event.String()).
			Msg("Guild event not handled, ignoring it")
		return
	}

	service.discordService.PublishWebhook(ctx.CorrelationID, service.internalWebhookID,
		service.internalWebhookToken, content)
}
