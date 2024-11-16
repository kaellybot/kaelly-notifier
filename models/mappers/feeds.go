package mappers

import (
	"time"

	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
)

func MapFeed(rssMessage *amqp.NewsRSSMessage) *discordgo.WebhookParams {
	return &discordgo.WebhookParams{
		Username:  constants.ExternalName,
		AvatarURL: constants.AvatarURL,
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: rssMessage.Title,
				URL:   rssMessage.Url,
				Color: constants.RSSColor,
				Author: &discordgo.MessageEmbedAuthor{
					Name: rssMessage.AuthorName,
					URL:  rssMessage.Url,
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: constants.RSSLogo,
				},
				Image: &discordgo.MessageEmbedImage{
					URL: rssMessage.IconUrl,
				},
				Timestamp: rssMessage.Date.AsTime().Format(time.RFC3339),
			},
		},
	}
}
