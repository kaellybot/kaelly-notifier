package notifiers

import amqp "github.com/kaellybot/kaelly-amqp"

const (
	newsQueueName  = "notifier-news"
	newsRoutingkey = "news.*"
)

type Service interface {
	Consume() error
}

type Impl struct {
	broker amqp.MessageBroker
	// TODO Discord webhook
	// TODO Repositories
}
