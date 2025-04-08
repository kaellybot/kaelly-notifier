package constants

import "github.com/rs/zerolog"

const (
	LogCorrelationID = "correlationID"
	LogChannelID     = "channelID"
	LogEmojiType     = "emojiType"
	LogEntityCount   = "entityCount"
	LogEntityID      = "entityID"
	LogEvent         = "event"
	LogFileName      = "fileName"
	LogGame          = "game"
	LogImageURL      = "imageURL"
	LogLocale        = "locale"
	LogPanic         = "panic"

	LogLevelFallback = zerolog.InfoLevel
)
