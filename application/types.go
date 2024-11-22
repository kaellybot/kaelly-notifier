package application

import (
	"github.com/go-co-op/gocron/v2"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/services/discord"
	"github.com/kaellybot/kaelly-notifier/services/notifiers"
	"github.com/kaellybot/kaelly-notifier/services/workers"
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
	scheduler       gocron.Scheduler
	probes          insights.Probes
	prom            insights.PrometheusMetrics
	discordService  discord.Service
	workerService   workers.Service
	notifierService notifiers.Service
}
