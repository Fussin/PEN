package recon

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type SubdomainScanScanner struct{}

func (s *SubdomainScanScanner) Init(config *common.ScannerConfig) {}

func (s *SubdomainScanScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *SubdomainScanScanner) Name() string {
	return "SubdomainScanScanner"
}
