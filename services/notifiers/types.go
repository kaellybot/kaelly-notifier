package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/repositories/almanaxes"
	"github.com/kaellybot/kaelly-notifier/repositories/feeds"
	"github.com/kaellybot/kaelly-notifier/repositories/twitch"
	"github.com/kaellybot/kaelly-notifier/repositories/youtube"
	"github.com/kaellybot/kaelly-notifier/services/discord"
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
	almanaxRepo          almanaxes.Repository
	feedRepo             feeds.Repository
	twitchRepo           twitch.Repository
	youtubeRepo          youtube.Repository
	internalWebhookID    string
	internalWebhookToken string
}
