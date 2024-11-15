package twitch

import (
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	Get(streamerID string) ([]*entities.WebhookTwitch, error)
	Save(channelWebhook entities.WebhookTwitch) error
	Delete(channelWebhook entities.WebhookTwitch) error
}

type Impl struct {
	db databases.MySQLConnection
}
