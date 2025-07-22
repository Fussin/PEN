package modules

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type WebsocketScanner struct{}

func (s *WebsocketScanner) Init(config *common.ScannerConfig) {}

func (s *WebsocketScanner) Scan(target string, ctx context.Context) ([]common.Finding, error) {
	return nil, nil
}

func (s *WebsocketScanner) Name() string {
	return "WebsocketScanner"
}
