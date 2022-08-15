package usecase

import (
	"context"
	"github.com/monopeelz/linear-avocado/internal/project/entity"
	"github.com/monopeelz/linear-avocado/internal/scanner/ports"
	"github.com/monopeelz/linear-avocado/pkg/scanner"
	"os"
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
	ports.ProjectExplorer
	ports.Downloader
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
	findings := make([]scanner.Finding, 0)
	dest, err := s.Download(i.URL)
	if err != nil {
		return nil, err
	}
	paths, err := s.Explore(dest)
	if err != nil {
		return nil, err
	}
	for _, sca := range s.scanners {
		for _, p := range paths {
			f, _ := sca.ScanFile(p)
			findings = append(findings, f...)
		}
	}
	return findings, err
}

func NewScannerUseCase(e ports.ProjectExplorer, d ports.Downloader) ScannerUseCase {
	return &scannerUseCase{
		ProjectExplorer: e,
		Downloader:      d,
	}
}
