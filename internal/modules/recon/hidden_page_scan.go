package recon

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type HiddenPageScanScanner struct{}

func (s *HiddenPageScanScanner) Init(config *common.ScannerConfig) {}

func (s *HiddenPageScanScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *HiddenPageScanScanner) Name() string {
	return "HiddenPageScanScanner"
}
