package file_access

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type LfiScanner struct{}

func (s *LfiScanner) Init(config *common.ScannerConfig) {}

func (s *LfiScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *LfiScanner) Name() string {
	return "LfiScanner"
}
