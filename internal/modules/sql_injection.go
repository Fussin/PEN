package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type SqlInjectionScanner struct{}

func (s *SqlInjectionScanner) Init(config *common.ScannerConfig) {}

func (s *SqlInjectionScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *SqlInjectionScanner) Name() string {
	return "SqlInjectionScanner"
}
