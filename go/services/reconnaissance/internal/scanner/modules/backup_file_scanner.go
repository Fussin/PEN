package modules

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/scanner"
)

type BackupFileScanner struct{}

func (s *BackupFileScanner) Scan(ctx context.Context, target *scanner.Target, vulnChan chan<- *scanner.Vulnerability) {
	// Placeholder implementation
}
