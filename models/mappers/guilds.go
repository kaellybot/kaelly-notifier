package mappers

import (
	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func MapGuildCreateNews(guildNews *amqp.NewsGuildMessage,
	locale amqp.Language) *discordgo.WebhookParams {
	lg := constants.GetLanguage(locale)
	return &discordgo.WebhookParams{
		Content: i18n.Get(lg.Locale, "guild.create", i18n.Vars{
			"name":        guildNews.Name,
			"memberCount": guildNews.MemberCount,
		}),
	}
}

func MapGuildDeleteNews(guildNews *amqp.NewsGuildMessage,
	locale amqp.Language) *discordgo.WebhookParams {
	lg := constants.GetLanguage(locale)
	return &discordgo.WebhookParams{
		Content: i18n.Get(lg.Locale, "guild.delete", i18n.Vars{
			"name":        guildNews.Name,
			"memberCount": guildNews.MemberCount,
		}),
	}
}
