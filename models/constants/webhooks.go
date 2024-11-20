package constants

import "time"

const (
	Delta    = 2 * time.Hour
	MaxRetry = 5
)

type Webhook struct {
	WebhookID    string `gorm:"unique;not null"`
	WebhookToken string
	RetryNumber  int64 `gorm:"default:0"`
	// To update only in case of failure.
	PublishedAt *time.Time
	FailedAt    *time.Time
}
