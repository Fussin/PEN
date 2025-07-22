package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type GraphqlScanner struct{}

func (s *GraphqlScanner) Init(config *common.ScannerConfig) {}

func (s *GraphqlScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *GraphqlScanner) Name() string {
	return "GraphqlScanner"
}
