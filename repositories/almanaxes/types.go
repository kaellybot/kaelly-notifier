package almanaxes

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	Get(game amqp.Game, date time.Time) ([]*entities.WebhookAlmanax, error)
	Save(channelWebhook entities.WebhookAlmanax) error
	Delete(channelWebhook entities.WebhookAlmanax) error
}

type Impl struct {
	db databases.MySQLConnection
}
