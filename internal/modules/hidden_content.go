package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type HiddenContentScanner struct{}

func (s *HiddenContentScanner) Init(config *common.ScannerConfig) {}

func (s *HiddenContentScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *HiddenContentScanner) Name() string {
	return "HiddenContentScanner"
}
