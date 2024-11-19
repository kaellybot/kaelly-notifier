package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/kaellybot/kaelly-notifier/services/notifiers"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
	"github.com/kaellybot/kaelly-notifier/utils/insights"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	broker          amqp.MessageBroker
	db              databases.MySQLConnection
	probes          insights.Probes
	prom            insights.PrometheusMetrics
	discordService  discord.Service
	notifierService notifiers.Service
}
