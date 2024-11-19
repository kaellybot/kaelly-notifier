package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
)

type WebhookFeed struct {
	constants.Webhook
	FeedTypeID string    `gorm:"primaryKey"`
	Game       amqp.Game `gorm:"primaryKey"`
	Locale     amqp.Language
}
