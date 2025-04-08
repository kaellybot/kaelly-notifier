package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func MapSetNews(setNews *amqp.NewsSetMessage, game amqp.Game) string {
	return i18n.Get(constants.InternalLocale, "set.message", i18n.Vars{
		"game": constants.GetGame(game).Name,
		"sets": setNews.SetIds,
	})
}
