package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
)

type WebhookTwitter struct {
	constants.Webhook
	TwitterID string `gorm:"primaryKey"`
	Locale    amqp.Language
}
