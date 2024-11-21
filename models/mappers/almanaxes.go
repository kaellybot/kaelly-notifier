package mappers

import (
	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/services/emojis"
	"github.com/kaellybot/kaelly-notifier/utils/discord"
	"github.com/kaellybot/kaelly-notifier/utils/translators"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func MapAlmanax(almanax *amqp.NewsAlmanaxMessage_I18NAlmanax,
	source *amqp.Source, emojiService emojis.Service) *discordgo.WebhookParams {
	lg := constants.GetLanguage(almanax.Locale)
	season := constants.GetSeason(almanax.Almanax.Date.AsTime())
	fullDate := lg.DateTranslator.FmtDateFull(almanax.Almanax.Date.AsTime())
	simpleDate := almanax.Almanax.Date.AsTime().Format(constants.DiscordDateOnlyFormat)
	return &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: i18n.Get(lg.Locale, "almanax.title", i18n.Vars{"date": fullDate}),
				URL: i18n.Get(lg.Locale, "almanax.url", i18n.Vars{
					"date": almanax.Almanax.Date.AsTime().Format(constants.KrosmozAlmanaxDateFormat),
				}),
				Color:     season.Color,
				Thumbnail: &discordgo.MessageEmbedThumbnail{URL: season.AlmanaxIcon},
				Image:     &discordgo.MessageEmbedImage{URL: almanax.Almanax.Tribute.Item.Icon},
				Author: &discordgo.MessageEmbedAuthor{
					Name:    source.GetName(),
					URL:     source.GetUrl(),
					IconURL: source.GetIcon(),
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  i18n.Get(lg.Locale, "almanax.bonus.title"),
						Value: almanax.Almanax.Bonus,
					},
					{
						Name: i18n.Get(lg.Locale, "almanax.tribute.title"),
						Value: i18n.Get(lg.Locale, "almanax.tribute.description", i18n.Vars{
							"item":     almanax.Almanax.Tribute.Item.Name,
							"emoji":    emojiService.GetItemTypeStringEmoji(almanax.Almanax.GetTribute().Item.GetType()),
							"quantity": almanax.Almanax.Tribute.Quantity,
						}),
					},
					{
						Name: i18n.Get(lg.Locale, "almanax.reward.title"),
						Value: i18n.Get(lg.Locale, "almanax.reward.description", i18n.Vars{
							"reward":   translators.FormatNumber(almanax.Almanax.Reward, lg.Locale),
							"kamaIcon": emojiService.GetMiscStringEmoji(constants.EmojiIDKama),
						}),
					},
				},
				Footer: discord.BuildDefaultFooter(lg.Locale, simpleDate),
			},
		},
	}
}
