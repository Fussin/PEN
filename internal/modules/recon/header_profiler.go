package recon

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type HeaderProfilerScanner struct{}

func (s *HeaderProfilerScanner) Init(config *common.ScannerConfig) {}

func (s *HeaderProfilerScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *HeaderProfilerScanner) Name() string {
	return "HeaderProfilerScanner"
}
