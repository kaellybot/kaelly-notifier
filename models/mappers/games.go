package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/i18n"
	di18n "github.com/kaysoro/discordgo-i18n"
)

func MapGameNews(gameNews *amqp.NewsGameMessage, game amqp.Game) string {
	return di18n.Get(i18n.InternalLocale, "game.message", di18n.Vars{
		"game":    constants.GetGame(game).Name,
		"version": gameNews.Version,
	})
}
