package mappers

import (
	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func MapGameNews(gameNews *amqp.NewsGameMessage, game amqp.Game,
	locale amqp.Language) *discordgo.WebhookParams {
	lg := constants.GetLanguage(locale)
	return &discordgo.WebhookParams{
		Content: i18n.Get(lg.Locale, "game.message", i18n.Vars{
			"game":    constants.GetGame(game).Name,
			"version": gameNews.Version,
		}),
	}
}
