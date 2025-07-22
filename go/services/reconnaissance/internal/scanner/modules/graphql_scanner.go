package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type GraphqlScanner struct{}

func (s *GraphqlScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
