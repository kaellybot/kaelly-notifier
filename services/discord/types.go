package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Service interface {
	PublishWebhook(webhookID, webhookToken string,
		content *discordgo.WebhookParams) error
	Shutdown()
}

type Impl struct {
	session *discordgo.Session
}
