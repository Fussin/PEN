package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type HostHeaderScanner struct{}

func (s *HostHeaderScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
