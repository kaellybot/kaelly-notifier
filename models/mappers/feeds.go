package mappers

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/i18n"
	"github.com/kaellybot/kaelly-notifier/utils/discord"
	"github.com/kaellybot/kaelly-notifier/utils/images"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

var (
	newLinesRegex = regexp.MustCompile("\n\\s*\n")
)

func MapFeed(rssMessage *amqp.NewsRSSMessage, locale amqp.Language) *discordgo.MessageSend {
	imageEmbed, files := retrieveImage(rssMessage.IconUrl)
	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       rssMessage.Title,
				Description: mapHTMLToDiscordMarkdown(rssMessage.Description, constants.EmbedDescriptionLimit),
				Author: &discordgo.MessageEmbedAuthor{
					Name: rssMessage.AuthorName,
				},
				Color: constants.RSSColor,
				URL:   rssMessage.Url,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: constants.RSSLogo,
				},
				Image:     imageEmbed,
				Timestamp: rssMessage.Date.AsTime().Format(time.RFC3339),
				Footer:    discord.BuildDefaultFooter(i18n.GetLanguage(locale).Locale, ""),
			},
		},
		Files: files,
	}
}

func retrieveImage(url string) (*discordgo.MessageEmbedImage, []*discordgo.File) {
	var imageEmbed *discordgo.MessageEmbedImage
	var files []*discordgo.File

	if url != "" {
		filename := path.Base(url)
		buffer, errGetImg := images.GetImageFromURL(context.Background(), url)
		if errGetImg != nil {
			log.Warn().Err(errGetImg).
				Str(constants.LogImageURL, url).
				Msgf("Cannot retrieve image, continuing without it")
			return imageEmbed, files
		}

		imageEmbed = &discordgo.MessageEmbedImage{
			URL: fmt.Sprintf("attachment://%v", filename),
		}

		files = []*discordgo.File{
			{
				Name:   filename,
				Reader: bytes.NewReader(buffer.Bytes()),
			},
		}
	}

	return imageEmbed, files
}

func mapHTMLToDiscordMarkdown(input string, limit int) string {
	tokenizer := html.NewTokenizer(strings.NewReader(input))
	var output bytes.Buffer
	var listDepth int

	for {
		tt := tokenizer.Next()
		//nolint:exhaustive // No need for other cases.
		switch tt {
		case html.ErrorToken:
			cleaned := output.String()
			decoded := html.UnescapeString(cleaned)
			decoded = newLinesRegex.ReplaceAllString(decoded, "\n")

			if len(decoded) > limit {
				decoded = decoded[:limit]
				lastSpace := strings.LastIndex(decoded, " ")
				if lastSpace != -1 {
					decoded = decoded[:lastSpace]
				}
				decoded += "..."
			}

			return decoded

		case html.TextToken:
			output.Write(tokenizer.Text())

		case html.StartTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "h1":
				output.WriteString("\n# ")
			case "h2":
				output.WriteString("\n## ")
			case "h3":
				output.WriteString("\n### ")
			case "li":
				output.WriteString(strings.Repeat("  ", listDepth) + "* ")
			case "ul", "ol":
				listDepth++
			case "br", "p", "div":
				output.WriteString("\n")
			}

		case html.EndTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "ul", "ol":
				listDepth--
			case "h1", "h2", "h3", "p", "div", "li":
				output.WriteString("\n")
			}
		}
	}
}
