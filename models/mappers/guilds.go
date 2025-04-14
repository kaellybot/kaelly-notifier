package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/i18n"
	di18n "github.com/kaysoro/discordgo-i18n"
)

func MapGuildCreateNews(guildNews *amqp.NewsGuildMessage) string {
	return di18n.Get(i18n.InternalLocale, "guild.create", di18n.Vars{
		"name":        guildNews.Name,
		"memberCount": guildNews.MemberCount,
	})
}

func MapGuildDeleteNews(guildNews *amqp.NewsGuildMessage) string {
	return di18n.Get(i18n.InternalLocale, "guild.delete", di18n.Vars{
		"name":        guildNews.Name,
		"memberCount": guildNews.MemberCount,
	})
}
