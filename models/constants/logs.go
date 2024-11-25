package constants

import "github.com/rs/zerolog"

const (
	LogCorrelationID = "correlationID"
	LogEmojiType     = "emojiType"
	LogEntityCount   = "entityCount"
	LogEntityID      = "entityID"
	LogEvent         = "event"
	LogFileName      = "fileName"
	LogGame          = "game"
	LogImageURL      = "imageURL"
	LogLocale        = "locale"
	LogPanic         = "panic"
	LogWebhookCount  = "webhookCount"
	LogWebhookID     = "webhookID"

	LogLevelFallback = zerolog.InfoLevel
)
