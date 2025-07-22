package injection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type XpathScanner struct{}

func (s *XpathScanner) Init(config *common.ScannerConfig) {}

func (s *XpathScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *XpathScanner) Name() string {
	return "XpathScanner"
}
