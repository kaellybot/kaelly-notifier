package entities

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
)

type WebhookAlmanax struct {
	constants.Webhook
	Game   amqp.Game `gorm:"primaryKey"`
	Locale amqp.Language
}
