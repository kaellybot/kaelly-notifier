package constants

import (
	"github.com/rs/zerolog"
)

const (
	ConfigFileName = ".env"

	// Discord Bot Token.
	DiscordToken = "DISCORD_TOKEN"

	// Channel snowflake for internal informations.
	ReportingChannelID = "REPORTING_CHANNEL_ID"

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

	// Probe port.
	ProbePort = "PROBE_PORT"

	// Metric port.
	MetricPort = "METRIC_PORT"

	// Zerolog values from [trace, debug, info, warn, error, fatal, panic].
	LogLevel = "LOG_LEVEL"

	// Boolean; used to register commands at development guild level or globally.
	Production = "PRODUCTION"

	defaultDiscordToken       = ""
	defaultReportingChannelID = ""
	defaultMySQLURL           = "localhost:3306"
	defaultMySQLUser          = ""
	defaultMySQLPassword      = ""
	defaultMySQLDatabase      = "kaellybot"
	defaultRabbitMQAddress    = "amqp://localhost:5672"
	defaultProbePort          = 9090
	defaultMetricPort         = 2112
	defaultLogLevel           = zerolog.InfoLevel
	defaultProduction         = false
)

func GetDefaultConfigValues() map[string]any {
	return map[string]any{
		DiscordToken:       defaultDiscordToken,
		ReportingChannelID: defaultReportingChannelID,
		MySQLURL:           defaultMySQLURL,
		MySQLUser:          defaultMySQLUser,
		MySQLPassword:      defaultMySQLPassword,
		MySQLDatabase:      defaultMySQLDatabase,
		RabbitMQAddress:    defaultRabbitMQAddress,
		ProbePort:          defaultProbePort,
		MetricPort:         defaultMetricPort,
		LogLevel:           defaultLogLevel.String(),
		Production:         defaultProduction,
	}
}
