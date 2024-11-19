package constants

import "github.com/rs/zerolog"

const (
	LogCorrelationID         = "correlationID"
	LogWebhookCount          = "webhookCount"
	LogSucceededWebhookCount = "succeededWebhookCount"
	LogEntityID              = "entityID"
	LogFileName              = "fileName"
	LogGame                  = "game"
	LogLocale                = "locale"
	LogWebhookID             = "webhookID"

	LogLevelFallback = zerolog.InfoLevel
)
