package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type CrlfScanner struct{}

func (s *CrlfScanner) Init(config *common.ScannerConfig) {}

func (s *CrlfScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *CrlfScanner) Name() string {
	return "CrlfScanner"
}
