package constants

import (
	"github.com/rs/zerolog"
)

const (
	ConfigFileName = ".env"

	// Discord Bot Token.
	DiscordToken = "DISCORD_TOKEN"

	// Discord webhook ID used for internal purposes.
	DiscordWebhookID = "DISCORD_WEBHOOK_ID"

	//nolint:gosec // False positive.
	// Discord webhook token used for internal purposes.
	DiscordWebhookToken = "DISCORD_WEBHOOK_TOKEN"

	// MySQL URL with the following format: HOST:PORT.
	MySQLURL = "MYSQL_URL"

	// MySQL user.
	MySQLUser = "MYSQL_USER"

	// MySQL password.
	MySQLPassword = "MYSQL_PASSWORD"

	// MySQL database name.
	MySQLDatabase = "MYSQL_DATABASE"

	// RabbitMQ address.
	RabbitMQAddress = "RABBITMQ_ADDRESS"

	// Metric port.
	MetricPort = "METRIC_PORT"

	// Zerolog values from [trace, debug, info, warn, error, fatal, panic].
	LogLevel = "LOG_LEVEL"

	// Boolean; used to register commands at development guild level or globally.
	Production = "PRODUCTION"

	defaultDiscordToken        = ""
	defaultDiscordWebhookID    = ""
	defaultDiscordWebhookToken = ""
	defaultMySQLURL            = "localhost:3306"
	defaultMySQLUser           = ""
	defaultMySQLPassword       = ""
	defaultMySQLDatabase       = "kaellybot"
	defaultRabbitMQAddress     = "amqp://localhost:5672"
	defaultMetricPort          = 2112
	defaultLogLevel            = zerolog.InfoLevel
	defaultProduction          = false
)

func GetDefaultConfigValues() map[string]any {
	return map[string]any{
		DiscordToken:        defaultDiscordToken,
		DiscordWebhookID:    defaultDiscordWebhookID,
		DiscordWebhookToken: defaultDiscordWebhookToken,
		MySQLURL:            defaultMySQLURL,
		MySQLUser:           defaultMySQLUser,
		MySQLPassword:       defaultMySQLPassword,
		MySQLDatabase:       defaultMySQLDatabase,
		RabbitMQAddress:     defaultRabbitMQAddress,
		MetricPort:          defaultMetricPort,
		LogLevel:            defaultLogLevel.String(),
		Production:          defaultProduction,
	}
}
