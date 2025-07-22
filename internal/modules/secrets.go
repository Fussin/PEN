package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type SecretsScanner struct{}

func (s *SecretsScanner) Init(config *common.ScannerConfig) {}

func (s *SecretsScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *SecretsScanner) Name() string {
	return "SecretsScanner"
}
