package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	dg, err := discordgo.New(fmt.Sprintf("Bot %v", viper.GetString(constants.DiscordToken)))
	if err != nil {
		log.Error().Err(err).Msgf("Connecting to Discord gateway failed")
		return nil, err
	}

	return &Impl{
		session: dg,
	}, nil
}

func (service *Impl) AnnounceMessage(correlationID, newsChannelID string,
	message *discordgo.MessageSend) {
	msg, errSend := service.session.ChannelMessageSendComplex(newsChannelID, message)
	if errSend != nil {
		log.Error().Err(errSend).
			Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogChannelID, newsChannelID).
			Msgf("Cannot send message in news channel, ignoring it")
		return
	}

	_, errCrossPost := service.session.ChannelMessageCrosspost(newsChannelID, msg.ID)
	if errCrossPost != nil {
		log.Error().Err(errCrossPost).
			Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogChannelID, newsChannelID).
			Msgf("Cannot crosspost message in news channel, ignoring it")
		return
	}
}

func (service *Impl) SendMessage(correlationID, channelID, content string) {
	_, errSend := service.session.ChannelMessageSend(channelID, content)
	if errSend != nil {
		log.Error().Err(errSend).
			Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogChannelID, channelID).
			Msgf("Cannot send message in channel, ignoring it")
	}
}

func (service *Impl) Shutdown() {
	log.Info().Msgf("Closing Discord connections...")
	err := service.session.Close()
	if err != nil {
		log.Warn().Err(err).Msgf("Cannot close session and shutdown correctly")
	}
}
