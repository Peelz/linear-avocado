package usecase

import (
	"context"
	"fmt"
	"github.com/monopeelz/linear-avocado/internal/project/entity"
	"github.com/monopeelz/linear-avocado/pkg/scanner"
	"os"
	"path/filepath"
)

// ScannerUseCase use for register scanner.Scanner and exec it all.
type ScannerUseCase interface {
	AddScanner(scanner scanner.Scanner)
	Rules() []scanner.Rule
	Exec(ctx context.Context, i entity.Project) ([]scanner.Finding, error)
}

type scannerUseCase struct {
	files    []os.File
	scanned  []os.File
	scanners []scanner.Scanner
	ProjectExplorer
}

func (s scannerUseCase) AddScanner(scanner scanner.Scanner) {
	// TODO implement me
	panic("implement me")
}

func (s scannerUseCase) Rules() []scanner.Rule {
	r := make([]scanner.Rule, 0)
	for _, sca := range s.scanners {
		r = append(r, sca.Rules()...)
	}
	return r
}

func (s scannerUseCase) Exec(ctx context.Context, i entity.Project) ([]scanner.Finding, error) {

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
}

func NewScannerUseCase(e ProjectExplorer) ScannerUseCase {
	return &scannerUseCase{
		ProjectExplorer: e,
	}
}
