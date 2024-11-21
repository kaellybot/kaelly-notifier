package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/repositories/webhooks"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/kaellybot/kaelly-notifier/services/emojis"
)

const (
	newsQueueName  = "notifier-news"
	newsRoutingkey = "news.*"
)

type Service interface {
	Consume()
}

type Impl struct {
	broker               amqp.MessageBroker
	discordService       discord.Service
	emojiService         emojis.Service
	webhookRepo          webhooks.Repository
	internalWebhookID    string
	internalWebhookToken string
}
