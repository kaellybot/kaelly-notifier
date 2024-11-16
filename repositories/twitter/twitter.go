package twitter

import (
	"time"

	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) Get(twitterID string, date time.Time) ([]*entities.WebhookTwitter, error) {
	var webhooks []*entities.WebhookTwitter
	err := repo.db.GetDB().
		Where("twitter_id = ? AND updated_at < ?", twitterID, date).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) Save(webhook entities.WebhookTwitter) error {
	return repo.db.GetDB().Save(&webhook).Error
}

func (repo *Impl) Delete(webhook entities.WebhookTwitter) error {
	if webhook != (entities.WebhookTwitter{}) {
		return repo.db.GetDB().Delete(&webhook).Error
	}

	return nil
}
