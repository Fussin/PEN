package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type XpathInjectionScanner struct{}

func (s *XpathInjectionScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
