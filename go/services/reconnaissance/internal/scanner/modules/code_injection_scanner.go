package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type CodeInjectionScanner struct{}

func (s *CodeInjectionScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
