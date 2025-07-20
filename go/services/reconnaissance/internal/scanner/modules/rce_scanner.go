package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type RceScanner struct{}

func (s *RceScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
