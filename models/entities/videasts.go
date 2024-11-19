package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
)

type WebhookYoutube struct {
	constants.Webhook
	VideastID string `gorm:"primaryKey"`
	Locale    amqp.Language
}
