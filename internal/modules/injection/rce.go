package injection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type RceScanner struct{}

func (s *RceScanner) Init(config *common.ScannerConfig) {}

func (s *RceScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *RceScanner) Name() string {
	return "RceScanner"
}
