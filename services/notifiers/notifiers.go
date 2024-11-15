package notifiers

import amqp "github.com/kaellybot/kaelly-amqp"

func New(broker amqp.MessageBroker) *Impl {
	return &Impl{
		broker: broker,
	}
}

func GetBinding() amqp.Binding {
	return amqp.Binding{
		Exchange:   amqp.ExchangeNews,
		RoutingKey: newsRoutingkey,
		Queue:      newsQueueName,
	}
}

func (service *Impl) Consume() error {
	// TODO
	return nil
}
