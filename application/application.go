package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/repositories/almanaxes"
	emojiRepo "github.com/kaellybot/kaelly-notifier/repositories/emojis"
	"github.com/kaellybot/kaelly-notifier/repositories/feeds"
	"github.com/kaellybot/kaelly-notifier/repositories/twitter"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/kaellybot/kaelly-notifier/services/emojis"
	"github.com/kaellybot/kaelly-notifier/services/news"
	"github.com/kaellybot/kaelly-notifier/services/notifiers"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
	"github.com/kaellybot/kaelly-notifier/utils/insights"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	broker := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress),
		amqp.WithBindings(notifiers.GetBinding()))

	db := databases.New()
	if err := db.Run(); err != nil {
		return nil, err
	}

	probes := insights.NewProbes(broker.IsConnected, db.IsConnected)
	prom := insights.NewPrometheusMetrics()

	// Repositories
	almanaxRepo := almanaxes.New(db)
	feedRepo := feeds.New(db)
	twitterRepo := twitter.New(db)
	emojiRepo := emojiRepo.New(db)

	// services
	discordService, errDisc := discord.New()
	if errDisc != nil {
		return nil, errDisc
	}

	emojiService, errEmoji := emojis.New(emojiRepo)
	if errEmoji != nil {
		return nil, errEmoji
	}

	newsService, errNew := news.New(almanaxRepo, feedRepo, twitterRepo)
	if errNew != nil {
		return nil, errNew
	}

	notifierService := notifiers.New(broker, discordService,
		emojiService, newsService)

	return &Impl{
		broker:          broker,
		db:              db,
		probes:          probes,
		prom:            prom,
		discordService:  discordService,
		notifierService: notifierService,
	}, nil
}

func (app *Impl) Run() error {
	app.probes.ListenAndServe()
	app.prom.ListenAndServe()

	if err := app.broker.Run(); err != nil {
		return err
	}

	app.notifierService.Consume()
	return nil
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	app.db.Shutdown()
	app.discordService.Shutdown()
	app.prom.Shutdown()
	app.probes.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
