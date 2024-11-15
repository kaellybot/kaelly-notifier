package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/services/notifiers"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	db              databases.MySQLConnection
	broker          amqp.MessageBroker
	notifierService notifiers.Service
}
