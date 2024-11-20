package constants

type Webhook struct {
	WebhookID    string `gorm:"unique;not null"`
	WebhookToken string
}
