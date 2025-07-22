package injection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type CodeInjectionScanner struct{}

func (s *CodeInjectionScanner) Init(config *common.ScannerConfig) {}

func (s *CodeInjectionScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *CodeInjectionScanner) Name() string {
	return "CodeInjectionScanner"
}
