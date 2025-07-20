package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type ClickjackingScanner struct{}

func (s *ClickjackingScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
