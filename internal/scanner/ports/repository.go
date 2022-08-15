package ports

import (
	"context"
	"github.com/monopeelz/linear-avocado/internal/scanner/models"
)

type ScannerRepository interface {
	UpdateJobOnProcess(ctx context.Context, i string) error
	UpdateJobOnSuccess(ctx context.Context, job models.Job) error
}
