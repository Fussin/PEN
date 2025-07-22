package injection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type NosqlInjectionScanner struct{}

func (s *NosqlInjectionScanner) Init(config *common.ScannerConfig) {}

func (s *NosqlInjectionScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *NosqlInjectionScanner) Name() string {
	return "NosqlInjectionScanner"
}
