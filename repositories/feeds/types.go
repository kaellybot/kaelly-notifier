package feeds

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	Get(feedTypeID string, game amqp.Game, locale amqp.Language,
		date time.Time) ([]*entities.WebhookFeed, error)
	Save(channelWebhook entities.WebhookFeed) error
	Delete(channelWebhook entities.WebhookFeed) error
}

type Impl struct {
	db databases.MySQLConnection
}
