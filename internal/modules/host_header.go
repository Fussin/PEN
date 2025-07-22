package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type HostHeaderScanner struct{}

func (s *HostHeaderScanner) Init(config *common.ScannerConfig) {}

func (s *HostHeaderScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *HostHeaderScanner) Name() string {
	return "HostHeaderScanner"
}
