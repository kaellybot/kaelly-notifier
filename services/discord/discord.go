package discord

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/services/workers"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New(workerService workers.Service) (*Impl, error) {
	dg, err := discordgo.New(fmt.Sprintf("Bot %v", viper.GetString(constants.DiscordToken)))
	if err != nil {
		log.Error().Err(err).Msgf("Connecting to Discord gateway failed")
		return nil, err
	}

	return &Impl{
		session:       dg,
		workerService: workerService,
	}, nil
}

func (service *Impl) PublishWebhook(correlationID, webhookID, webhookToken string,
	content *discordgo.WebhookParams) {
	service.workerService.Work(func() {
		_, err := service.session.WebhookExecute(
			webhookID,
			webhookToken,
			false, // No need to wait for webhook response.
			content,
		)
		if err != nil {
			log.Error().Err(err).
				Str(constants.LogCorrelationID, correlationID).
				Msgf("Cannot publish through webhook, ignoring it")
		}
	})
}

func (service *Impl) IsWebhookAvailable(webhookID string) bool {
	_, err := service.session.Webhook(webhookID)
	if err != nil {
		var httpErr *discordgo.RESTError
		if errors.As(err, &httpErr) && httpErr.Response.StatusCode == http.StatusNotFound {
			return false
		}
		log.Warn().Err(err).
			Str(constants.LogWebhookID, webhookID).
			Msg("Ignoring it this time...")
		return true
	}

	return true
}

func (service *Impl) Shutdown() {
	log.Info().Msgf("Closing Discord connections...")
	err := service.session.Close()
	if err != nil {
		log.Warn().Err(err).Msgf("Cannot close session and shutdown correctly")
	}
}
