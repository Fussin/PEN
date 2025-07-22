package file_access

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type DirTraversalScanner struct{}

func (s *DirTraversalScanner) Init(config *common.ScannerConfig) {}

func (s *DirTraversalScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *DirTraversalScanner) Name() string {
	return "DirTraversalScanner"
}
