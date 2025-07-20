package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type SstiScanner struct{}

func (s *SstiScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
