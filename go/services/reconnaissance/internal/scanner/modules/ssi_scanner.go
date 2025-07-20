package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type SsiScanner struct{}

func (s *SsiScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
