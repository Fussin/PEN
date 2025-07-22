package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type SsiScanner struct{}

func (s *SsiScanner) Init(config *common.ScannerConfig) {}

func (s *SsiScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *SsiScanner) Name() string {
	return "SsiScanner"
}
