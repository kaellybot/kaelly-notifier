package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) feedNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	feedSource := service.newsService.GetFeedSource(message.NewsRSSMessage.Type,
		message.Language, message.Game)
	if feedSource == nil {
		log.Error().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogEntityID, message.NewsRSSMessage.Type).
			Str(constants.LogGame, message.Game.String()).
			Str(constants.LogLocale, message.Language.String()).
			Msg("Cannot retrieve feed source, ignoring the feed occurence")
		return
	}

	embeds := mappers.MapFeed(message.NewsRSSMessage, message.Language)
	service.discordService.
		AnnounceMessage(ctx.CorrelationID, feedSource.NewsChannelID, embeds)
	log.Info().
		Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogEntityID, feedSource.FeedTypeID).
		Str(constants.LogGame, message.Game.String()).
		Str(constants.LogLocale, message.Language.String()).
		Msg("Feed published!")
}
