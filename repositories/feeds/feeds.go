package feeds

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(feedTypeID string, game amqp.Game, locale amqp.Language,
	date time.Time) ([]*entities.WebhookFeed, error) {
	var webhooks []*entities.WebhookFeed
	err := repo.db.GetDB().
		Where("feed_type_id = ? AND game = ? AND locale = ? AND updated_at < ?",
			feedTypeID, game, locale, date).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) Save(webhook entities.WebhookFeed) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookFeed) error {
	if webhook != (entities.WebhookFeed{}) {
		return repo.db.GetDB().Delete(&webhook).Error
	}

	return nil
}
