package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/mappers"
	"github.com/rs/zerolog/log"
)

func (service *Impl) almanaxNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	almanaxes := withFallback(message.NewsAlmanaxMessage.Almanaxes)
	for _, almanax := range almanaxes {
		almanaxNews := service.newsService.GetAlmanaxNews(almanax.Locale, message.Game)
		if almanaxNews == nil {
			log.Error().
				Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGame, message.Game.String()).
				Str(constants.LogLocale, almanax.Locale.String()).
				Msg("Cannot get almanax news, ignoring the occurence")
			continue
		}

		response := mappers.MapAlmanax(almanax, message.NewsAlmanaxMessage.Source,
			service.emojiService)
		service.discordService.AnnounceMessage(ctx.CorrelationID, almanaxNews.NewsChannelID, response)
		log.Info().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogGame, message.Game.String()).
			Str(constants.LogLocale, almanax.Locale.String()).
			Msg("Almanax published!")
	}
}

// Ensure missing Almanax for specific language are always set with fallback language.
// If fallback is missing too, language is ignored.
func withFallback(almanaxes []*amqp.NewsAlmanaxMessage_I18NAlmanax) []*amqp.NewsAlmanaxMessage_I18NAlmanax {
	var fallback *amqp.Almanax
	for _, almanax := range almanaxes {
		if almanax.Almanax != nil && almanax.Locale == constants.DefaultAMQPLocale {
			fallback = almanax.Almanax
			break
		}
	}

	result := make([]*amqp.NewsAlmanaxMessage_I18NAlmanax, 0)
	for _, almanax := range almanaxes {
		if almanax.Almanax != nil {
			result = append(result, &amqp.NewsAlmanaxMessage_I18NAlmanax{
				Almanax: almanax.Almanax,
				Locale:  almanax.Locale,
			})
		} else {
			if fallback != nil {
				log.Warn().
					Str(constants.LogLocale, almanax.Locale.String()).
					Msg("Almanax not available in this language, switching on fallback")

				result = append(result, &amqp.NewsAlmanaxMessage_I18NAlmanax{
					Almanax: fallback,
					Locale:  almanax.Locale,
				})
			} else {
				log.Error().
					Str(constants.LogLocale, almanax.Locale.String()).
					Msg("Cannot deliver almanax, fallback was not set")
			}
		}
	}

	return result
}
