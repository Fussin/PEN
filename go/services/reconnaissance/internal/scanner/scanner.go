package scanner

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/common"
)

type ScanResult struct {
	Target          *common.Target
	Vulnerabilities []*common.Vulnerability
}

type VulnerabilityScanner interface {
	Scan(ctx context.Context, target *common.Target, vulnChan chan<- *common.Vulnerability)
}
