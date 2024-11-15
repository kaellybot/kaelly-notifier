package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookYoutube struct {
	WebhookID    string
	WebhookToken string
	VideastID    string `gorm:"primaryKey"`
	Locale       amqp.Language
	RetryNumber  int64 `gorm:"default:0"`
	UpdatedAt    time.Time
}
