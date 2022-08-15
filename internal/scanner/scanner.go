package scanner

import (
	"database/sql"
	"github.com/monopeelz/linear-avocado/internal/scanner/adapter"
	"github.com/monopeelz/linear-avocado/internal/scanner/usecase"
	"go.uber.org/zap"
)

func Initial(db *sql.DB, logger *zap.Logger) *adapter.Handler {
	exp := adapter.NewProjectExplore()
	dw := adapter.NewGitAdapter()
	rp := adapter.NewScannerRepository(db, logger)
	u := usecase.NewScannerUseCase(exp, dw, rp, logger)

	u.AddScanner(adapter.NewSimpleScanner())
	return adapter.NewHandler(u)
}
