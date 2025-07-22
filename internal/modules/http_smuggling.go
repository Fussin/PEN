package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type HttpSmugglingScanner struct{}

func (s *HttpSmugglingScanner) Init(config *common.ScannerConfig) {}

func (s *HttpSmugglingScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *HttpSmugglingScanner) Name() string {
	return "HttpSmugglingScanner"
}
