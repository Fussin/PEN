package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type SqlScanner struct{}

func (s *SqlScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
