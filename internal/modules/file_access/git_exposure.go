package file_access

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type GitExposureScanner struct{}

func (s *GitExposureScanner) Init(config *common.ScannerConfig) {}

func (s *GitExposureScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *GitExposureScanner) Name() string {
	return "GitExposureScanner"
}
