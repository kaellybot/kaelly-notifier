package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Service interface {
	AnnounceMessage(correlationID, newsChannelID string,
		message *discordgo.MessageSend)
	SendMessage(correlationID, channelID, content string)
	Shutdown()
}

type Impl struct {
	session *discordgo.Session
}
