package constants

import "time"

const (
	Delta    = 2 * time.Hour
	MaxRetry = 10
)

type Webhook struct {
	WebhookID    string
	WebhookToken string
	RetryNumber  int64 `gorm:"default:0"`
	// To update only in case of failure.
	PublishedAt time.Time `gorm:"default:'0001-01-01 00:00:00';autoCreateTime:false;autoUpdateTime:false"`
	FailedAt    time.Time `gorm:"default:'0001-01-01 00:00:00';autoCreateTime:false;autoUpdateTime:false"`
}
