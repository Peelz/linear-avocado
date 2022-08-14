package project

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/monopeelz/linear-avocado/internal/project/adapter"
	"github.com/monopeelz/linear-avocado/internal/project/usecase"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func Initial(e *gin.Engine, db *sql.DB, amqpConn *amqp.Connection, logger *zap.Logger) {
	h := adapter.NewHandler(usecase.NewProjectUseCase(
		adapter.NewProjectRepository(db, logger),
		adapter.NewMessageBroker(amqpConn),
	))
	adapter.RegisterRoute(e, h)
}
