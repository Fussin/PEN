package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type IdorScanner struct{}

func (s *IdorScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
