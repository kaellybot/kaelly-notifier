package notifiers

import (
	"time"

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

func (service *Impl) dispatch(content *discordgo.WebhookParams, webhookModel any,
	webhooks []*constants.Webhook) int {
	var dispatched int
	updatedWebhooks := make([]*constants.Webhook, 0)
	excludedWebhooks := make([]*constants.Webhook, 0)

	// Try dispatching content through webhooks.
	for _, webhook := range webhooks {
		errPub := service.discordService.
			PublishWebhook(webhook.WebhookID, webhook.WebhookToken, content)

		if errPub != nil {
			if toKeep, updatedWebhook := service.applyFailurePolicy(webhook); toKeep {
				webhook = updatedWebhook
				updatedWebhooks = append(updatedWebhooks, webhook)
			} else {
				excludedWebhooks = append(excludedWebhooks, webhook)
			}
		} else {
			dispatched++
			updatedWebhooks = append(updatedWebhooks, service.applySuccessPolicy(webhook))
		}
	}

	// Updating webhooks which failed and errored webhooks which succeed this time.
	errUpdate := service.webhookRepo.UpdateWebhooks(webhookModel, updatedWebhooks)
	if errUpdate != nil {
		log.Error().Err(errUpdate).
			Msgf("Cannot update webhooks, ignoring them for this time")
	}

	errDel := service.webhookRepo.DeleteWebhooks(webhookModel, excludedWebhooks)
	if errDel != nil {
		log.Error().Err(errDel).
			Msgf("Cannot remove unreachable webhooks, ignoring them for this time")
	}

	return dispatched
}

func (service *Impl) applySuccessPolicy(webhook *constants.Webhook) *constants.Webhook {
	now := time.Now()
	webhook.PublishedAt = &now
	if webhook.RetryNumber != 0 {
		webhook.RetryNumber = 0
		return webhook
	}
	return webhook
}

func (service *Impl) applyFailurePolicy(webhook *constants.Webhook) (bool, *constants.Webhook) {
	now := time.Now()
	if webhook.FailedAt == nil || webhook.FailedAt.Add(constants.Delta).Before(time.Now()) {
		webhook.RetryNumber++
		webhook.FailedAt = &now
		if webhook.RetryNumber >= constants.MaxRetry {
			return false, webhook
		}
	}
	return true, webhook
}
