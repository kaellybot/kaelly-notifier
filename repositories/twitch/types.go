package twitch

import (
	"time"

	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	Get(streamerID string, date time.Time) ([]*entities.WebhookTwitch, error)
	Save(channelWebhook entities.WebhookTwitch) error
	Delete(channelWebhook entities.WebhookTwitch) error
}

type Impl struct {
	db databases.MySQLConnection
}
