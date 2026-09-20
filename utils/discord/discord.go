package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

// BuildDefaultFooter brands the embed with the bot of the news' game: each game has
// its own Discord application, so the footer must name the bot that actually posts it.
func BuildDefaultFooter(game constants.AnkamaGame, lg discordgo.Locale,
	date string) *discordgo.MessageEmbedFooter {
	return &discordgo.MessageEmbedFooter{
		Text: i18n.Get(lg, "default.footer", i18n.Vars{
			"name": game.BotName,
			"date": date,
		}),
		IconURL: game.BotAvatarURL,
	}
}
