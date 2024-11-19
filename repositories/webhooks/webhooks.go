package webhooks

import (
	"errors"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) GetAlmanaxWebhooks(game amqp.Game, locale amqp.Language,
) ([]*entities.WebhookAlmanax, error) {
	var webhooks []*entities.WebhookAlmanax
	err := repo.db.GetDB().
		Where("game = ? AND locale = ?", game, locale).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) GetFeedWebhooks(feedTypeID string, game amqp.Game, locale amqp.Language,
) ([]*entities.WebhookFeed, error) {
	var webhooks []*entities.WebhookFeed
	err := repo.db.GetDB().
		Where("feed_type_id = ? AND game = ? AND locale = ?",
			feedTypeID, game, locale).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) GetTwitchWebhooks(streamerID string) ([]*entities.WebhookTwitch, error) {
	var webhooks []*entities.WebhookTwitch
	err := repo.db.GetDB().
		Where("streamer_id = ?", streamerID).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) GetTwitterWebhooks(twitterID string) ([]*entities.WebhookTwitter, error) {
	var webhooks []*entities.WebhookTwitter
	err := repo.db.GetDB().
		Where("twitter_id = ?", twitterID).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) GetYoutubeWebhooks(videastID string) ([]*entities.WebhookYoutube, error) {
	var webhooks []*entities.WebhookYoutube
	err := repo.db.GetDB().
		Where("videast_id = ?", videastID).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) UpdateWebhooks(model any, webhooks []*constants.Webhook) error {
	var err error
	for _, wh := range webhooks {
		webhook := wh
		err = errors.Join(err, repo.db.GetDB().Model(model).Updates(webhook).Error)
	} // TODO
	return err
}

func (repo *Impl) DeleteWebhooks(model any, webhooks []*constants.Webhook) error {
	var err error
	for _, wh := range webhooks {
		webhook := wh
		err = errors.Join(err, repo.db.GetDB().Model(model).Delete(webhook).Error)
	} // TODO
	return err
}
