package main

import (
	"context"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/monopeelz/linear-avocado/internal/pkg/message"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type Config struct {
	DbUrl  string `env:"DB_URL"`
	AmqUrl string `env:"AMQ_URL"`
}

func main() {
	// Load .env file
	if os.Getenv("ENV_FILE") != "none" {
		godotenv.Load()
	}

	cfg := Config{}
	env.Parse(&cfg)

	conn, err := amqp.Dial(cfg.AmqUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	sb := message.NewRabbitMqSubscriber(ch)
	err = sb.Subscribe("scan-project", func(ctx context.Context, b []byte) {
		log.Printf("Received a message: %s", b)
	})
	if err != nil {
		failOnError(err, "Failed to subscribe")
	}
}
