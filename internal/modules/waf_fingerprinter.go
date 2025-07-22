package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type WafFingerprinterScanner struct{}

func (s *WafFingerprinterScanner) Init(config *common.ScannerConfig) {}

func (s *WafFingerprinterScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *WafFingerprinterScanner) Name() string {
	return "WafFingerprinterScanner"
}
