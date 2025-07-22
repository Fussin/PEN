package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type SecretScanner struct{}

func (s *SecretScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
