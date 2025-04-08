package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func MapGameNews(gameNews *amqp.NewsGameMessage, game amqp.Game) string {
	return i18n.Get(constants.InternalLocale, "game.message", i18n.Vars{
		"game":    constants.GetGame(game).Name,
		"version": gameNews.Version,
	})
}
