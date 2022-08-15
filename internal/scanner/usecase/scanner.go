package usecase

import (
	"context"
	"github.com/monopeelz/linear-avocado/internal/scanner/entity"
	"github.com/monopeelz/linear-avocado/internal/scanner/ports"
	"github.com/monopeelz/linear-avocado/pkg/scanner"
	"go.uber.org/zap"
)

// ScannerUseCase use for register scanner.Scanner and exec it all.
type ScannerUseCase interface {
	AddScanner(scanner scanner.Scanner)
	Scanners() []scanner.Scanner
	Rules() []scanner.Rule
	Exec(ctx context.Context, i entity.Project) ([]scanner.Finding, error)
	Files() []string
}

type scannerUseCase struct {
	files      []string
	scanners   []scanner.Scanner
	explorer   ports.ProjectExplorer
	downloader ports.Downloader
	rp         ports.ScannerRepository
	logger     *zap.Logger
}

func (s *scannerUseCase) Files() []string {
	return s.files
}

func (s *scannerUseCase) AddScanner(scanner scanner.Scanner) {
	s.scanners = append(s.scanners, scanner)
}

func (s *scannerUseCase) Scanners() []scanner.Scanner {
	return s.scanners
}

func (s *scannerUseCase) Rules() []scanner.Rule {
	r := make([]scanner.Rule, 0)
	for _, sca := range s.scanners {
		r = append(r, sca.Rules()...)
	}
	return r
}

func (s *scannerUseCase) Exec(ctx context.Context, i entity.Project) ([]scanner.Finding, error) {
	findings := make([]scanner.Finding, 0)
	dest, err := s.downloader.Download(i.URL)
	if err != nil {
		s.logger.Error("", zap.Error(err))
		return nil, err
	}
	paths, err := s.explorer.Explore(dest)
	if err != nil {
		s.logger.Error("", zap.Error(err))
		return nil, err
	}
	for _, sca := range s.scanners {
		for _, p := range paths {
			f, err := sca.ScanFile(p)
			if err != nil {
				s.logger.Error("", zap.Error(err))
			}
			if f != nil && len(f) > 0 {
				s.logger.Debug("Found", zap.Any("count", len(f)))
				findings = append(findings, f...)
			}
		}
	}
	return findings, err
}

func NewScannerUseCase(
	e ports.ProjectExplorer,
	d ports.Downloader,
	rp ports.ScannerRepository,
	logger *zap.Logger,
) ScannerUseCase {
	return &scannerUseCase{
		explorer:   e,
		downloader: d,
		rp:         rp,
		logger:     logger,
	}
}
