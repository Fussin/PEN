package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type ConfigDiscoveryScanner struct{}

func (s *ConfigDiscoveryScanner) Init(config *common.ScannerConfig) {}

func (s *ConfigDiscoveryScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *ConfigDiscoveryScanner) Name() string {
	return "ConfigDiscoveryScanner"
}
