package frontend

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type ApiScannerScanner struct{}

func (s *ApiScannerScanner) Init(config *common.ScannerConfig) {}

func (s *ApiScannerScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *ApiScannerScanner) Name() string {
	return "ApiScannerScanner"
}
