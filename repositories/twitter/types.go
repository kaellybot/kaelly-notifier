package twitter

import (
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	Get(twitterID string) ([]*entities.WebhookTwitter, error)
	Save(channelWebhook entities.WebhookTwitter) error
	Delete(channelWebhook entities.WebhookTwitter) error
}

type Impl struct {
	db databases.MySQLConnection
}