package scanner

import (
	"context"
)

type Target struct {
	URL string
}

type Vulnerability struct {
	Type     string
	Evidence string
}

type ScanResult struct {
	Target          *Target
	Vulnerabilities []*Vulnerability
}

type VulnerabilityScanner interface {
	Scan(ctx context.Context, target *Target, vulnChan chan<- *Vulnerability)
}
