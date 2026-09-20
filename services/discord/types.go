package discord

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
)

// Every game served by a bot needs its own token: starting without one would leave
// that game's news silently undelivered.
var errMissingDiscordToken = errors.New("no Discord token configured for a served game")

type Service interface {
	AnnounceMessage(correlationID string, game amqp.Game, newsChannelID string,
		message *discordgo.MessageSend)
	SendMessage(correlationID string, game amqp.Game, channelID, content string)
	Shutdown()
}

// Impl holds one session per enabled game. Each news channel belongs to one bot and
// only a message's author may crosspost it, so the session must always match the
// game of the channel being posted to.
type Impl struct {
	sessions map[amqp.Game]*discordgo.Session
}
