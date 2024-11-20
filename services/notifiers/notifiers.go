package notifiers

import (
	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/repositories/webhooks"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New(broker amqp.MessageBroker, discordService discord.Service,
	webhookRepo webhooks.Repository) *Impl {
	return &Impl{
		broker:               broker,
		discordService:       discordService,
		webhookRepo:          webhookRepo,
		internalWebhookID:    viper.GetString(constants.DiscordWebhookID),
		internalWebhookToken: viper.GetString(constants.DiscordWebhookToken),
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

func (service *Impl) dispatch(content *discordgo.WebhookParams, webhooks []*constants.Webhook) int {
	var dispatched int
	for _, webhook := range webhooks {
		errPub := service.discordService.
			PublishWebhook(webhook.WebhookID, webhook.WebhookToken, content)
		if errPub != nil {
			log.Debug().Err(errPub).
				Msgf("Could not publish webhook, continuing...")
			continue
		}

		dispatched++
	}

	return dispatched
}
