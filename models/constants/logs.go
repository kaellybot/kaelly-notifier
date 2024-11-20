package constants

import "github.com/rs/zerolog"

const (
	LogCorrelationID         = "correlationID"
	LogWebhookCount          = "webhookCount"
	LogSucceededWebhookCount = "succeededWebhookCount"
	LogEntityCount           = "entityCount"
	LogEntityID              = "entityID"
	LogFileName              = "fileName"
	LogGame                  = "game"
	LogImageURL              = "imageURL"
	LogLocale                = "locale"
	LogWebhookID             = "webhookID"

	LogLevelFallback = zerolog.InfoLevel
)
