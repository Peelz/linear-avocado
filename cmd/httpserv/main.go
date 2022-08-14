package main

import (
	"database/sql"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/monopeelz/linear-avocado/internal/project"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
	"log"
	"os"
)

type Config struct {
	Port   int    `env:"PORT" envDefault:"8080"`
	DbUrl  string `env:"DB_URL"`
	AmqUrl string `env:"AMQ_URL"`
}

func main() {
	// Load .env file
	if os.Getenv("ENV_FILE") != "none" {
		godotenv.Load()
	}

	// Init application stack dependencies
	cfg := Config{}
	env.Parse(&cfg)
	logger, _ := zap.NewDevelopment()
	e := gin.Default()
	// DB
	dbConn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DbUrl)))
	// AMQ
	amqConn, err := amqp.Dial(cfg.AmqUrl)
	if err != nil {
		panic(err)
	}
	// Dependencies injection
	project.Initial(e, dbConn, amqConn, logger)

	// Start http server
	log.Fatal(e.Run(fmt.Sprintf(":%v", cfg.Port)))
}
