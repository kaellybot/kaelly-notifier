package mappers

import (
	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
)

func MapAlmanax(_ *amqp.NewsAlmanaxMessage_I18NAlmanax) *discordgo.WebhookParams {
	return &discordgo.WebhookParams{
		// TODO
	}
}
