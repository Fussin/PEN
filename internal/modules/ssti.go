package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type SstiScanner struct{}

func (s *SstiScanner) Init(config *common.ScannerConfig) {}

func (s *SstiScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *SstiScanner) Name() string {
	return "SstiScanner"
}
