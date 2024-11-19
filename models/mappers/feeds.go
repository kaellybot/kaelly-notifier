package mappers

import (
	"time"

	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/utils/discord"
)

func MapFeed(rssMessage *amqp.NewsRSSMessage, locale amqp.Language) *discordgo.WebhookParams {
	return &discordgo.WebhookParams{
		Username:  constants.ExternalName,
		AvatarURL: constants.AvatarURL,
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: rssMessage.Title,
				Author: &discordgo.MessageEmbedAuthor{
					Name: rssMessage.AuthorName,
				},
				Color: constants.RSSColor,
				URL:   rssMessage.Url,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: constants.RSSLogo,
				},
				Image: &discordgo.MessageEmbedImage{
					URL: rssMessage.IconUrl, // TODO Image must be downloaded
				},
				Timestamp: rssMessage.Date.AsTime().Format(time.RFC3339),
				Footer:    discord.BuildDefaultFooter(constants.MapAMQPLocale(locale)),
			},
		},
	}
}
