package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type DirectoryTraversalScanner struct{}

func (s *DirectoryTraversalScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
