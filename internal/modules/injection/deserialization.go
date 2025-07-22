package injection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type DeserializationScanner struct{}

func (s *DeserializationScanner) Init(config *common.ScannerConfig) {}

func (s *DeserializationScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *DeserializationScanner) Name() string {
	return "DeserializationScanner"
}
