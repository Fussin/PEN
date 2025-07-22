package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type SsrfScanner struct{}

func (s *SsrfScanner) Init(config *common.ScannerConfig) {}

func (s *SsrfScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *SsrfScanner) Name() string {
	return "SsrfScanner"
}
