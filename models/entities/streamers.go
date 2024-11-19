package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
)

type WebhookTwitch struct {
	constants.Webhook
	StreamerID string `gorm:"primaryKey"`
	Locale     amqp.Language
}
