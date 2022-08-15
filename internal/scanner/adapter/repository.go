package adapter

import (
	"context"
	"database/sql"
	"github.com/monopeelz/linear-avocado/internal/scanner/models"
	"github.com/monopeelz/linear-avocado/internal/scanner/ports"
	"github.com/monopeelz/linear-avocado/pkg/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"go.uber.org/zap"
	"time"
)

type scannerRepository struct {
	db     *sql.DB
	bun    *bun.DB
	logger *zap.Logger
}

func (s scannerRepository) UpdateJobOnSuccess(ctx context.Context, job models.Job) error {
	s.logger.Debug("", zap.Any("id", job.UUID))
	_, err := s.bun.NewUpdate().
		Model(&job).
		Column("status", "finished_at", "detail").
		Where("uuid = ?", job.UUID).
		Exec(ctx)
	return utils.PgError(err)
}

func (s scannerRepository) UpdateJobOnProcess(ctx context.Context, i string) error {
	s.logger.Debug("", zap.Any("id", i))
	t := time.Now()
	_, err := s.bun.NewUpdate().
		Model(&models.Job{
			Status:    models.InProgress,
			StartedAt: &t,
		}).
		Column("status", "started_at").
		Where("uuid = ?", i).
		Exec(ctx)
	return utils.PgError(err)
}

func NewScannerRepository(db *sql.DB, logger *zap.Logger) ports.ScannerRepository {
	bunx := bun.NewDB(db, pgdialect.New())
	return &scannerRepository{
		db,
		bunx,
		logger,
	}
}
