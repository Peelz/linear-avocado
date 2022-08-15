package adapter

import (
	"database/sql"
	"github.com/monopeelz/linear-avocado/internal/scanner/ports"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"go.uber.org/zap"
)

type scannerRepository struct {
	db     *sql.DB
	bun    *bun.DB
	logger *zap.Logger
}

func NewScannerRepository(db *sql.DB, logger *zap.Logger) ports.ScannerRepository {
	bunx := bun.NewDB(db, pgdialect.New())
	return &scannerRepository{
		db,
		bunx,
		logger,
	}
}
