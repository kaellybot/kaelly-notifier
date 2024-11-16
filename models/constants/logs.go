package constants

import "github.com/rs/zerolog"

const (
	LogCorrelationID         = "correlationID"
	LogWebhookCount          = "webhookCount"
	LogSucceededWebhookCount = "succeededWebhookCount"
	LogFileName              = "fileName"
	LogLocale                = "locale"

	LogLevelFallback = zerolog.InfoLevel
)
