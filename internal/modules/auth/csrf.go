package auth

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type CsrfScanner struct{}

func (s *CsrfScanner) Init(config *common.ScannerConfig) {}

func (s *CsrfScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *CsrfScanner) Name() string {
	return "CsrfScanner"
}
