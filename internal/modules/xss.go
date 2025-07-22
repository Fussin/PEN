package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type XssScanner struct{}

func (s *XssScanner) Init(config *common.ScannerConfig) {}

func (s *XssScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *XssScanner) Name() string {
	return "XssScanner"
}
