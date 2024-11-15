package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookAlmanax struct {
	WebhookID    string
	WebhookToken string
	Game         amqp.Game `gorm:"primaryKey"`
	Locale       amqp.Language
	RetryNumber  int64 `gorm:"default:0"`
	UpdatedAt    time.Time
}
