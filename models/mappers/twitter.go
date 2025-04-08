package mappers

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/utils/discord"
)

func MapTweet(tweet *amqp.NewsTwitterMessage, locale amqp.Language) *discordgo.MessageSend {
	description := tweet.Description
	if len(description) > constants.EmbedDescriptionLimit {
		description = fmt.Sprintf("%v...", description[:constants.EmbedDescriptionLimit])
	}

	firstEmbed := &discordgo.MessageEmbed{
		Title:       tweet.Title,
		Description: description,
		Color:       constants.TwitterColor,
		URL:         tweet.Url,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: constants.TwitterLogo,
		},
		Timestamp: tweet.Date.AsTime().Format(time.RFC3339),
		Footer:    discord.BuildDefaultFooter(constants.GetLanguage(locale).Locale, ""),
	}

	if len(tweet.IconUrls) > 0 {
		firstEmbed.Image = &discordgo.MessageEmbedImage{
			URL: tweet.IconUrls[0],
		}
	}
	embeds := []*discordgo.MessageEmbed{
		firstEmbed,
	}

	for i := 1; i < len(tweet.IconUrls); i++ {
		embeds = append(embeds, &discordgo.MessageEmbed{
			URL: tweet.Url,
			Image: &discordgo.MessageEmbedImage{
				URL: tweet.IconUrls[i],
			},
		})
	}

	return &discordgo.MessageSend{
		Embeds: embeds,
	}
}
