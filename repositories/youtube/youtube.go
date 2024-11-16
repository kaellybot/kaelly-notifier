package youtube

import (
	"time"

	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(videastID string, date time.Time) ([]*entities.WebhookYoutube, error) {
	var webhooks []*entities.WebhookYoutube
	err := repo.db.GetDB().
		Where("videast_id = ? AND updated_at < ?", videastID, date).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) Save(webhook entities.WebhookYoutube) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookYoutube) error {
	if webhook != (entities.WebhookYoutube{}) {
		return repo.db.GetDB().Delete(&webhook).Error
	}

	return nil
}
