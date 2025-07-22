package auth

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type CorsScanner struct{}

func (s *CorsScanner) Init(config *common.ScannerConfig) {}

func (s *CorsScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *CorsScanner) Name() string {
	return "CorsScanner"
}
