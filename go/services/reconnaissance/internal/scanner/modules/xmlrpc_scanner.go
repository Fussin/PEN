package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type XmlrpcScanner struct{}

func (s *XmlrpcScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
