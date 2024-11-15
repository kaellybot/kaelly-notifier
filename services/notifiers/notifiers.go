package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/repositories/almanaxes"
	"github.com/kaellybot/kaelly-notifier/repositories/feeds"
	"github.com/kaellybot/kaelly-notifier/repositories/twitch"
	"github.com/kaellybot/kaelly-notifier/repositories/youtube"
	"github.com/rs/zerolog/log"
)

func New(broker amqp.MessageBroker, almanaxRepo almanaxes.Repository, feedRepo feeds.Repository,
	twitchRepo twitch.Repository, youtubeRepo youtube.Repository) *Impl {
	return &Impl{
		broker:      broker,
		almanaxRepo: almanaxRepo,
		feedRepo:    feedRepo,
		twitchRepo:  twitchRepo,
		youtubeRepo: youtubeRepo,
	}
}

func GetBinding() amqp.Binding {
	return amqp.Binding{
		Exchange:   amqp.ExchangeNews,
		RoutingKey: newsRoutingkey,
		Queue:      newsQueueName,
	}
}

func (service *Impl) Consume() {
	log.Info().Msgf("Consuming news...")
	service.broker.Consume(newsQueueName, service.consume)
}

func (service *Impl) consume(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	//exhaustive:ignore Don't need to be exhaustive here since they will be handled by default case
	switch message.Type {
	case amqp.RabbitMQMessage_NEWS_ALMANAX:
		service.almanaxNews(ctx, message)
	case amqp.RabbitMQMessage_NEWS_GAME:
		service.gameNews(ctx, message)
	case amqp.RabbitMQMessage_NEWS_RSS:
		service.feedNews(ctx, message)
	case amqp.RabbitMQMessage_NEWS_SET:
		service.setNews(ctx, message)
	case amqp.RabbitMQMessage_NEWS_TWITTER:
		service.twitterNews(ctx, message)
	default:
		log.Warn().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Type not recognized, request ignored")
	}
}
