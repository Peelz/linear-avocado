package adapter

import (
	"context"
	"github.com/monopeelz/linear-avocado/internal/project/ports"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBroker struct {
	conn *amqp.Connection
}

func (msg MessageBroker) Publish(ctx context.Context, topic string, b []byte) error {
	var err error

	ch, err := msg.conn.Channel()
	q, err := ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})
	return err
}

func NewMessageBroker(conn *amqp.Connection) ports.MessagePublisher {
	return &MessageBroker{
		conn,
	}
}
