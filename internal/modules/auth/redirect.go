package auth

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type RedirectScanner struct{}

func (s *RedirectScanner) Init(config *common.ScannerConfig) {}

func (s *RedirectScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *RedirectScanner) Name() string {
	return "RedirectScanner"
}
