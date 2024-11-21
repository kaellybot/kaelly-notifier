package notifiers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron/v2"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/repositories/webhooks"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/kaellybot/kaelly-notifier/services/emojis"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New(broker amqp.MessageBroker, scheduler gocron.Scheduler,
	discordService discord.Service, emojiService emojis.Service,
	webhookRepo webhooks.Repository) (*Impl, error) {
	service := Impl{
		broker:               broker,
		discordService:       discordService,
		emojiService:         emojiService,
		webhookRepo:          webhookRepo,
		internalWebhookID:    viper.GetString(constants.DiscordWebhookID),
		internalWebhookToken: viper.GetString(constants.DiscordWebhookToken),
	}
	_, errJob := scheduler.NewJob(
		gocron.CronJob(viper.GetString(constants.WebhookPurgeCronTab), true),
		gocron.NewTask(func() { service.purgeWebhooks() }),
		gocron.WithName("Purge unused webhooks"),
	)
	if errJob != nil {
		return nil, errJob
	}

	return &service, nil
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

func (service *Impl) purgeWebhooks() {
	log.Info().Msgf("Purging unused webhooks...")

	models := []any{
		&entities.WebhookAlmanax{},
		&entities.WebhookFeed{},
		&entities.WebhookTwitch{},
		&entities.WebhookTwitter{},
		&entities.WebhookYoutube{},
	}

	var purged int
	for _, model := range models {
		purged += service.purgeWebhookByTypes(model)
	}

	log.Info().
		Int(constants.LogEntityCount, purged).
		Msg("Webhooks purged!")
}

func (service *Impl) purgeWebhookByTypes(model any) int {
	webhookIDs, errGet := service.webhookRepo.GetWebhookIDs(model)
	if errGet != nil {
		log.Warn().Err(errGet).
			Msg("Cannot retrieve webhooks from DB, ignoring...")
		return 0
	}

	purgedWebhookIDs := make([]string, 0)
	for _, webhookID := range webhookIDs {
		if !service.discordService.IsWebhookAvailable(webhookID) {
			purgedWebhookIDs = append(purgedWebhookIDs, webhookID)
		}
	}

	errDel := service.webhookRepo.DeleteWebhooks(purgedWebhookIDs, model)
	if errDel != nil {
		log.Warn().Err(errDel).
			Msg("Cannot delete unused webhooks from DB, ignoring...")
		return 0
	}

	return len(purgedWebhookIDs)
}
