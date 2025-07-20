package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type NosqlScanner struct{}

func (s *NosqlScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
