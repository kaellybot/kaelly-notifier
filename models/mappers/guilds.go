package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/i18n"
	di18n "github.com/kaysoro/discordgo-i18n"
)

// The report carries no game: it is posted by the bot of its game, so the message
// author already identifies which bot gained or lost the guild.
func MapGuildCreateNews(guildNews *amqp.NewsGuildMessage) string {
	return di18n.Get(i18n.InternalLocale, "guild.create", di18n.Vars{
		"name":        guildNews.GetName(),
		"memberCount": guildNews.GetMemberCount(),
	})
}

func MapGuildDeleteNews(guildNews *amqp.NewsGuildMessage) string {
	return di18n.Get(i18n.InternalLocale, "guild.delete", di18n.Vars{
		"name":        guildNews.GetName(),
		"memberCount": guildNews.GetMemberCount(),
	})
}
