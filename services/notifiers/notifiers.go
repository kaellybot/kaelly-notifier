package notifiers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/kaellybot/kaelly-notifier/services/emojis"
	"github.com/kaellybot/kaelly-notifier/services/news"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New(broker amqp.MessageBroker, discordService discord.Service,
	emojiService emojis.Service, newsService news.Service) *Impl {
	return &Impl{
		reportingChannelID: viper.GetString(constants.ReportingChannelID),
		broker:             broker,
		discordService:     discordService,
		emojiService:       emojiService,
		newsService:        newsService,
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
	case amqp.RabbitMQMessage_NEWS_GUILD:
		service.guildNews(ctx, message)
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
