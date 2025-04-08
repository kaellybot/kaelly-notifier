package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/kaellybot/kaelly-notifier/services/emojis"
	"github.com/kaellybot/kaelly-notifier/services/news"
)

const (
	newsQueueName  = "notifier-news"
	newsRoutingkey = "news.*"
)

type Service interface {
	Consume()
}

type Impl struct {
	reportingChannelID string
	broker             amqp.MessageBroker
	discordService     discord.Service
	emojiService       emojis.Service
	newsService        news.Service
}
