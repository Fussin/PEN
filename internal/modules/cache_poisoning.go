package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type CachePoisoningScanner struct{}

func (s *CachePoisoningScanner) Init(config *common.ScannerConfig) {}

func (s *CachePoisoningScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *CachePoisoningScanner) Name() string {
	return "CachePoisoningScanner"
}
