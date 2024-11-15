package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookTwitch struct {
	WebhookID    string
	WebhookToken string
	StreamerID   string `gorm:"primaryKey"`
	Locale       amqp.Language
	RetryNumber  int64 `gorm:"default:0"`
	UpdatedAt    time.Time
}
