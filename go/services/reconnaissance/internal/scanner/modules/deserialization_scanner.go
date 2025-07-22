package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type DeserializationScanner struct{}

func (s *DeserializationScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
