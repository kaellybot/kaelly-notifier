package constants

import "time"

const (
	Delta    = 2 * time.Hour
	MaxRetry = 10
)

type Webhook struct {
	WebhookID    string
	WebhookToken string
	RetryNumber  int64     `gorm:"default:0"`
	UpdatedAt    time.Time // To update only in case of failure.
}
