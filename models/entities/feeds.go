package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookFeed struct {
	WebhookID    string
	WebhookToken string
	FeedTypeID   string    `gorm:"primaryKey"`
	Game         amqp.Game `gorm:"primaryKey"`
	Locale       amqp.Language
	RetryNumber  int64 `gorm:"default:0"`
	UpdatedAt    time.Time
}
