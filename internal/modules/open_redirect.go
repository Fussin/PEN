package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type OpenRedirectScanner struct{}

func (s *OpenRedirectScanner) Init(config *common.ScannerConfig) {}

func (s *OpenRedirectScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *OpenRedirectScanner) Name() string {
	return "OpenRedirectScanner"
}
