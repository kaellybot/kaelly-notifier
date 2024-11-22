package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-notifier/services/workers"
)

type Service interface {
	PublishWebhook(correlationID, webhookID, webhookToken string,
		content *discordgo.WebhookParams)
	IsWebhookAvailable(webhookID string) bool
	Shutdown()
}

type Impl struct {
	session       *discordgo.Session
	workerService workers.Service
}
