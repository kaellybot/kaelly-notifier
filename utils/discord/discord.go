package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func BuildDefaultFooter(lg discordgo.Locale, date string) *discordgo.MessageEmbedFooter {
	return &discordgo.MessageEmbedFooter{
		Text: i18n.Get(lg, "default.footer", i18n.Vars{
			"name": constants.ExternalName,
			"date": date,
		}),
		IconURL: constants.AvatarURL,
	}
}
