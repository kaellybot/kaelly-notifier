package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/repositories/almanaxes"
	"github.com/kaellybot/kaelly-notifier/repositories/feeds"
	"github.com/kaellybot/kaelly-notifier/repositories/twitch"
	"github.com/kaellybot/kaelly-notifier/repositories/youtube"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/kaellybot/kaelly-notifier/services/notifiers"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	db, errDB := databases.New()
	if errDB != nil {
		log.Fatal().Err(errDB).Msgf("DB instantiation failed, shutting down.")
	}

	broker := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress),
		amqp.WithBindings(notifiers.GetBinding()))

	// Repositories
	almanaxRepo := almanaxes.New(db)
	feedRepo := feeds.New(db)
	twitchRepo := twitch.New(db)
	youtubeRepo := youtube.New(db)

	// services
	discordService, errDisc := discord.New(viper.GetString(constants.DiscordToken))
	if errDisc != nil {
		log.Fatal().Err(errDB).Msgf("Discord connection failed, shutting down.")
	}

	notifierService := notifiers.New(broker, discordService, almanaxRepo,
		feedRepo, twitchRepo, youtubeRepo)

	return &Impl{
		db:              db,
		broker:          broker,
		notifierService: notifierService,
	}, nil
}

func (app *Impl) Run() error {
	errBroker := app.broker.Run()
	if errBroker != nil {
		return errBroker
	}

	app.notifierService.Consume()
	return nil
}

func (app *Impl) Shutdown() {
	app.db.Shutdown()
	app.broker.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
