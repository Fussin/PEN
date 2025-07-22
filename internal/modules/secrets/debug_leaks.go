package secrets

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type DebugLeaksScanner struct{}

func (s *DebugLeaksScanner) Init(config *common.ScannerConfig) {}

func (s *DebugLeaksScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *DebugLeaksScanner) Name() string {
	return "DebugLeaksScanner"
}
