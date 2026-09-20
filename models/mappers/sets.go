package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/i18n"
	di18n "github.com/kaysoro/discordgo-i18n"
)

func MapSetNews(setNews *amqp.NewsSetMessage, game constants.AnkamaGame) string {
	return di18n.Get(i18n.InternalLocale, "set.message", di18n.Vars{
		"userID":        constants.AuthorID,
		"game":          game.Name,
		"createdSetIDs": len(setNews.GetCreatedSetIds()),
		"updatedSetIDs": len(setNews.GetUpdatedSetIds()),
		"deletedSetIDs": len(setNews.GetDeletedSetIds()),
	})
}
