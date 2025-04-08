package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) twitterNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	twitterAccount := service.newsService.GetTwitterAccount(message.NewsTwitterMessage.TwitterId)
	if twitterAccount == nil {
		log.Error().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogEntityID, message.NewsTwitterMessage.TwitterId).
			Str(constants.LogGame, message.Game.String()).
			Str(constants.LogLocale, message.Language.String()).
			Msg("Cannot retrieve twitter account, ignoring the tweet occurence")
		return
	}

	response := mappers.MapTweet(message.NewsTwitterMessage, message.Language)
	service.discordService.
		AnnounceMessage(ctx.CorrelationID, twitterAccount.NewsChannelID, response)
	log.Info().
		Str(constants.LogCorrelationID, ctx.CorrelationID).
		Str(constants.LogChannelID, twitterAccount.NewsChannelID).
		Str(constants.LogEntityID, twitterAccount.ID).
		Str(constants.LogGame, message.Game.String()).
		Str(constants.LogLocale, message.Language.String()).
		Msg("Tweet published!")
}
