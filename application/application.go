package application

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	emojiRepo "github.com/kaellybot/kaelly-notifier/repositories/emojis"
	"github.com/kaellybot/kaelly-notifier/repositories/webhooks"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/kaellybot/kaelly-notifier/services/emojis"
	"github.com/kaellybot/kaelly-notifier/services/notifiers"
	"github.com/kaellybot/kaelly-notifier/services/workers"
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

	scheduler, errScheduler := gocron.NewScheduler(gocron.WithLocation(time.UTC))
	if errScheduler != nil {
		return nil, errScheduler
	}

	// Repositories
	webhooksRepo := webhooks.New(db)
	emojiRepo := emojiRepo.New(db)

	// services
	workerService := workers.New()
	discordService, errDisc := discord.New(workerService)
	if errDisc != nil {
		return nil, errDisc
	}

	emojiService, errEmoji := emojis.New(emojiRepo)
	if errEmoji != nil {
		return nil, errEmoji
	}

	notifierService, errNotif := notifiers.New(broker, scheduler, discordService,
		emojiService, webhooksRepo)
	if errNotif != nil {
		return nil, errNotif
	}

	return &Impl{
		broker:          broker,
		db:              db,
		scheduler:       scheduler,
		probes:          probes,
		prom:            prom,
		discordService:  discordService,
		workerService:   workerService,
		notifierService: notifierService,
	}, nil
}

func (app *Impl) Run() error {
	app.probes.ListenAndServe()
	app.prom.ListenAndServe()

	if err := app.broker.Run(); err != nil {
		return err
	}

	app.scheduler.Start()
	for _, job := range app.scheduler.Jobs() {
		scheduledTime, err := job.NextRun()
		if err == nil {
			log.Info().Msgf("%v scheduled at %v", job.Name(), scheduledTime)
		}
	}

	app.notifierService.Consume()
	return nil
}

func (app *Impl) Shutdown() {
	if err := app.scheduler.Shutdown(); err != nil {
		log.Error().Err(err).Msg("Cannot shutdown scheduler, continuing...")
	}

	app.broker.Shutdown()
	app.db.Shutdown()
	app.workerService.Shutdown()
	app.discordService.Shutdown()
	app.prom.Shutdown()
	app.probes.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
