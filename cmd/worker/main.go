package main

import (
	"database/sql"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/monopeelz/linear-avocado/internal/pkg/message"
	"github.com/monopeelz/linear-avocado/internal/scanner"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
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

	// Logger
	logger, _ := zap.NewDevelopment()

	// DB
	dbConn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DbUrl)))
	h := scanner.Initial(dbConn, logger)

	sb := message.NewRabbitMqSubscriber(ch)
	err = sb.Subscribe("scan-project", h.ScanProject)
	if err != nil {
		failOnError(err, "Failed to subscribe")
	}
}
