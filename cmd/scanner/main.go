package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/engine"
	"github.com/autonomouspen/scanner/internal/reporting"
)

func main() {
	target := flag.String("target", "", "The target URL to scan.")
	configFile := flag.String("config", "config.yml", "The path to the configuration file.")
	reportFile := flag.String("report", "report.json", "The path to the report file.")
	reportFormat := flag.String("format", "json", "The format of the report (json, csv, html).")
	enabledScanners := flag.String("scanners", "all", "A comma-separated list of scanners to enable.")
	flag.Parse()

	if *target == "" {
		log.Fatal("Target URL is required.")
	}

	// In a real implementation, we would load the config from the file.
	config := &common.ScannerConfig{
		MaxWorkers:         10,
		RequestTimeout:     10 * time.Second,
		EnableScreenshots:  true,
		RateLimit:          100,
		RateLimitInterval:  time.Second,
		RateLimitMaxTokens: 100,
		PayloadsFile:       "payloads.json",
	}

	if *enabledScanners != "all" {
		config.EnabledScanners = strings.Split(*enabledScanners, ",")
	}

	e := engine.NewEngine(config)
	findings, err := e.Scan(*target)
	if err != nil {
		log.Fatalf("Error scanning target: %v", err)
	}

	reporter := reporting.NewReporter()
	if err := reporter.Generate(findings, *reportFormat, *reportFile); err != nil {
		log.Fatalf("Error generating report: %v", err)
	}

	fmt.Printf("Scan complete. Report saved to %s\n", *reportFile)
}
