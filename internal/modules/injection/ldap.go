package injection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type LdapScanner struct{}

func (s *LdapScanner) Init(config *common.ScannerConfig) {}

func (s *LdapScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *LdapScanner) Name() string {
	return "LdapScanner"
}
