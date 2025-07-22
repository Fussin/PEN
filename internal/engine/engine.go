package engine

import (
	"context"
	"log"
	"sync"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/modules/auth"
	"github.com/autonomouspen/scanner/internal/modules/file_access"
	"github.com/autonomouspen/scanner/internal/modules/frontend"
	"github.com/autonomouspen/scanner/internal/modules/injection"
	"github.com/autonomouspen/scanner/internal/modules/recon"
	"github.com/autonomouspen/scanner/internal/modules/secrets"
)

type Engine struct {
	scanners []common.Scanner
	config   *common.ScannerConfig
}

func NewEngine(config *common.ScannerConfig) *Engine {
	engine := &Engine{
		config: config,
	}
	engine.registerScanners()
	return engine
}

func (e *Engine) registerScanners() {
	// This is where we would register all the scanners.
	// For now, we'll just register a few as an example.
	e.scanners = append(e.scanners, &injection.SQLInjectionScanner{})
	e.scanners = append(e.scanners, &frontend.XSSScanner{})
}

func (e *Engine) Scan(target string) ([]common.Finding, error) {
	var findings []common.Finding
	var wg sync.WaitGroup
	findingChan := make(chan []common.Finding, len(e.scanners))

	ctx := context.Background()

	for _, scanner := range e.scanners {
		wg.Add(1)
		go func(s common.Scanner) {
			defer wg.Done()
			s.Init(e.config)
			fs, err := s.Scan(target, ctx)
			if err != nil {
				log.Printf("Error scanning with %s: %v", s.Name(), err)
				return
			}
			findingChan <- fs
		}(scanner)
	}

	go func() {
		wg.Wait()
		close(findingChan)
	}()

	for fs := range findingChan {
		findings = append(findings, fs...)
	}

	return findings, nil
}
