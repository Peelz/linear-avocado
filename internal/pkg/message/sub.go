package message

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

// HandlerFunc callback function for subscribe
type HandlerFunc func(ctx context.Context, b []byte)

// Subscriber Message queue subscriber interface.
// Used for subscribe topic and handle topic with HandlerFunc
type Subscriber interface {
	Subscribe(topic string, handle HandlerFunc) error
}

type RabbitMqSubscriber struct {
	ch  *amqp.Channel
	cfg ChannelConfig
}

func (r RabbitMqSubscriber) Subscribe(topic string, handle HandlerFunc) error {

	q, err := r.ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}
	msgs, err := r.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}
	var forever chan time.Time
	go func() {
		for d := range msgs {
			handle(context.TODO(), d.Body)
		}
	}()
	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}

func NewRabbitMqSubscriber(ch *amqp.Channel) Subscriber {
	return &RabbitMqSubscriber{
		ch: ch,
	}
}
