package auth

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type CachePoisonScanner struct{}

func (s *CachePoisonScanner) Init(config *common.ScannerConfig) {}

func (s *CachePoisonScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *CachePoisonScanner) Name() string {
	return "CachePoisonScanner"
}
