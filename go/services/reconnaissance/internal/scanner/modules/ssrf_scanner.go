package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type SsrfScanner struct{}

func (s *SsrfScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
