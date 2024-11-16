package youtube

import (
	"time"

	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	Get(videastID string, date time.Time) ([]*entities.WebhookYoutube, error)
	Save(channelWebhook entities.WebhookYoutube) error
	Delete(channelWebhook entities.WebhookYoutube) error
}

type Impl struct {
	db databases.MySQLConnection
}
