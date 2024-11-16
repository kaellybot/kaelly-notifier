package twitch

import (
	"time"

	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(streamerID string, date time.Time) ([]*entities.WebhookTwitch, error) {
	var webhooks []*entities.WebhookTwitch
	err := repo.db.GetDB().
		Where("streamer_id = ? AND updated_at < ?", streamerID, date).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) Save(webhook entities.WebhookTwitch) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookTwitch) error {
	if webhook != (entities.WebhookTwitch{}) {
		return repo.db.GetDB().Delete(&webhook).Error
	}

	return nil
}
