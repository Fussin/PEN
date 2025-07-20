package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type WebsocketScanner struct{}

func (s *WebsocketScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
