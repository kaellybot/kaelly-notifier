package almanaxes

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(game amqp.Game, date time.Time) ([]*entities.WebhookAlmanax, error) {
	var webhooks []*entities.WebhookAlmanax
	err := repo.db.GetDB().
		Where("game = ? AND updated_at < ?", game, date).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) Save(webhook entities.WebhookAlmanax) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookAlmanax) error {
	if webhook != (entities.WebhookAlmanax{}) {
		return repo.db.GetDB().Delete(&webhook).Error
	}

	return nil
}
