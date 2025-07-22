package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type ClickjackingScanner struct{}

func (s *ClickjackingScanner) Init(config *common.ScannerConfig) {}

func (s *ClickjackingScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *ClickjackingScanner) Name() string {
	return "ClickjackingScanner"
}
