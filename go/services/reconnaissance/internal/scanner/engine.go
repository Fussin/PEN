package scanner

import (
	"context"
	"github.com/autonomouspen/reconnaissance/internal/common"
	"github.com/autonomouspen/reconnaissance/internal/scanner/modules"
	"sync"
	"time"
)

type ScanEngine struct {
	xssScanner *modules.XSSScanner
}

func NewScanEngine() *ScanEngine {
	return &ScanEngine{
		xssScanner: modules.NewXSSScanner(&modules.ScannerConfig{}),
	}
}

func (e *ScanEngine) ScanTarget(target *common.Target) (*ScanResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 24*time.Hour)
	defer cancel()

	result := &ScanResult{
		Target:          target,
		Vulnerabilities: make([]*common.Vulnerability, 0),
	}

	var wg sync.WaitGroup
	vulnChan := make(chan *common.Vulnerability, 1000)

	scanners := []VulnerabilityScanner{
		e.xssScanner,
	}

	for _, scanner := range scanners {
		wg.Add(1)
		go func(s VulnerabilityScanner) {
			defer wg.Done()
			s.Scan(ctx, target, vulnChan)
		}(scanner)
	}

	go func() {
		wg.Wait()
		close(vulnChan)
	}()

	for vuln := range vulnChan {
		result.Vulnerabilities = append(result.Vulnerabilities, vuln)
	}

	return result, nil
}
