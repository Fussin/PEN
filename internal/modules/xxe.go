package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type XxeScanner struct{}

func (s *XxeScanner) Init(config *common.ScannerConfig) {}

func (s *XxeScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *XxeScanner) Name() string {
	return "XxeScanner"
}
