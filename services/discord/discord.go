package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	sessions := make(map[amqp.Game]*discordgo.Session)
	for _, game := range constants.GetGames() {
		if game.DiscordTokenKey == "" {
			continue
		}

		token := viper.GetString(game.DiscordTokenKey)
		if token == "" {
			log.Error().Err(errMissingDiscordToken).
				Str(constants.LogGame, game.AMQPGame.String()).
				Msgf("%v is required to serve %v news", game.DiscordTokenKey, game.Name)
			return nil, fmt.Errorf("%w: %v", errMissingDiscordToken, game.DiscordTokenKey)
		}

		session, err := discordgo.New(fmt.Sprintf("Bot %v", token))
		if err != nil {
			log.Error().Err(err).
				Str(constants.LogGame, game.AMQPGame.String()).
				Msgf("Connecting to Discord gateway failed")
			return nil, err
		}

		sessions[game.AMQPGame] = session
		log.Info().
			Str(constants.LogGame, game.AMQPGame.String()).
			Msgf("Game enabled, news will be posted by %v", game.BotName)
	}

	return &Impl{sessions: sessions}, nil
}

func (service *Impl) AnnounceMessage(correlationID string, game amqp.Game, newsChannelID string,
	message *discordgo.MessageSend) {
	session, found := service.session(correlationID, game)
	if !found {
		return
	}

	msg, errSend := session.ChannelMessageSendComplex(newsChannelID, message)
	if errSend != nil {
		log.Error().Err(errSend).
			Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogGame, game.String()).
			Str(constants.LogChannelID, newsChannelID).
			Msgf("Cannot send message in news channel, ignoring it")
		return
	}

	// Crossposting uses the same session: only the author of a message may crosspost it.
	_, errCrossPost := session.ChannelMessageCrosspost(newsChannelID, msg.ID)
	if errCrossPost != nil {
		log.Error().Err(errCrossPost).
			Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogGame, game.String()).
			Str(constants.LogChannelID, newsChannelID).
			Msgf("Cannot crosspost message in news channel, ignoring it")
		return
	}
}

func (service *Impl) SendMessage(correlationID string, game amqp.Game, channelID, content string) {
	session, found := service.session(correlationID, game)
	if !found {
		return
	}

	_, errSend := session.ChannelMessageSend(channelID, content)
	if errSend != nil {
		log.Error().Err(errSend).
			Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogGame, game.String()).
			Str(constants.LogChannelID, channelID).
			Msgf("Cannot send message in channel, ignoring it")
	}
}

// session returns the bot that serves this game. Every game with a bot has a session
// by construction, so a miss means a game no bot serves, such as Retro. Its news is
// dropped rather than posted with another game's bot, which would put it in the wrong
// bot's name and make crossposting impossible.
func (service *Impl) session(correlationID string, game amqp.Game) (*discordgo.Session, bool) {
	session, found := service.sessions[game]
	if !found {
		log.Warn().
			Str(constants.LogCorrelationID, correlationID).
			Str(constants.LogGame, game.String()).
			Msgf("No bot serves this game, message dropped")
	}

	return session, found
}

func (service *Impl) Shutdown() {
	log.Info().Msgf("Closing Discord connections...")
	for game, session := range service.sessions {
		if err := session.Close(); err != nil {
			log.Warn().Err(err).
				Str(constants.LogGame, game.String()).
				Msgf("Cannot close session and shutdown correctly")
		}
	}
}
