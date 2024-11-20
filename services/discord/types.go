package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Service interface {
	PublishWebhook(webhookID, webhookToken string,
		content *discordgo.WebhookParams) error
	IsWebhookAvailable(webhookID string) bool
	Shutdown()
}

type Impl struct {
	session *discordgo.Session
}
