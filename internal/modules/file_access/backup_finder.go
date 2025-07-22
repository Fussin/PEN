package file_access

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type BackupFinderScanner struct{}

func (s *BackupFinderScanner) Init(config *common.ScannerConfig) {}

func (s *BackupFinderScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *BackupFinderScanner) Name() string {
	return "BackupFinderScanner"
}
