package auth

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type JwtScanner struct{}

func (s *JwtScanner) Init(config *common.ScannerConfig) {}

func (s *JwtScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *JwtScanner) Name() string {
	return "JwtScanner"
}
