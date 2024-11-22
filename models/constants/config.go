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

	// Worker pool number, dedicated to execute tasks like dispatch news.
	WorkerPool = "WORKER_POOL"

	// Task buffer size, dedicated to buffer tasks when all worker are busy.
	TaskBuffer = "TASK_BUFFER"

	// Cron tab to purge unused webhooks.
	WebhookPurgeCronTab = "WEBHOOK_PURGE_CRON_TAB"

	// Probe port.
	ProbePort = "PROBE_PORT"

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
	defaultWorkerPool          = 10
	defaultTaskBuffer          = 20
	defaultWebhookPurgeCronTab = "0 0 2 * * *"
	defaultProbePort           = 9090
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
		WorkerPool:          defaultWorkerPool,
		TaskBuffer:          defaultTaskBuffer,
		WebhookPurgeCronTab: defaultWebhookPurgeCronTab,
		ProbePort:           defaultProbePort,
		MetricPort:          defaultMetricPort,
		LogLevel:            defaultLogLevel.String(),
		Production:          defaultProduction,
	}
}
