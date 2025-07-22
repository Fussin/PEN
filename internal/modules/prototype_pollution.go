package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type PrototypePollutionScanner struct{}

func (s *PrototypePollutionScanner) Init(config *common.ScannerConfig) {}

func (s *PrototypePollutionScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *PrototypePollutionScanner) Name() string {
	return "PrototypePollutionScanner"
}
