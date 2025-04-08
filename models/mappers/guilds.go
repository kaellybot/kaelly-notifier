package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func MapGuildCreateNews(guildNews *amqp.NewsGuildMessage) string {
	return i18n.Get(constants.InternalLocale, "guild.create", i18n.Vars{
		"name":        guildNews.Name,
		"memberCount": guildNews.MemberCount,
	})
}

func MapGuildDeleteNews(guildNews *amqp.NewsGuildMessage) string {
	return i18n.Get(constants.InternalLocale, "guild.delete", i18n.Vars{
		"name":        guildNews.Name,
		"memberCount": guildNews.MemberCount,
	})
}
