package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type LfiScanner struct{}

func (s *LfiScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
