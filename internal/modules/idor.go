package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type IdorScanner struct{}

func (s *IdorScanner) Init(config *common.ScannerConfig) {}

func (s *IdorScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *IdorScanner) Name() string {
	return "IdorScanner"
}
