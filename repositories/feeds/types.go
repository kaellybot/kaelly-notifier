package feeds

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	Get(feedTypeID string, game amqp.Game) ([]*entities.WebhookFeed, error)
	Save(channelWebhook entities.WebhookFeed) error
	Delete(channelWebhook entities.WebhookFeed) error
}

type Impl struct {
	db databases.MySQLConnection
}
