package webhooks

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Repository interface {
	GetAlmanaxWebhooks(game amqp.Game, locale amqp.Language) ([]*entities.WebhookAlmanax, error)
	GetFeedWebhooks(feedTypeID string, game amqp.Game, locale amqp.Language,
	) ([]*entities.WebhookFeed, error)
	GetTwitchWebhooks(streamerID string) ([]*entities.WebhookTwitch, error)
	GetTwitterWebhooks(twitterID string) ([]*entities.WebhookTwitter, error)
	GetYoutubeWebhooks(videastID string) ([]*entities.WebhookYoutube, error)
	DeleteWebhooks(webhookIDs []string, model any) error
}

type Impl struct {
	db databases.MySQLConnection
}
