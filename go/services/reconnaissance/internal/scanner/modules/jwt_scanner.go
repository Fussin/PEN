package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type JwtScanner struct{}

func (s *JwtScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
