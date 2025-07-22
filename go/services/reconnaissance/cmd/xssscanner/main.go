package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/autonomouspen/reconnaissance/internal/common"
	"github.com/autonomouspen/reconnaissance/internal/reports"
	"github.com/autonomouspen/reconnaissance/internal/scanner/modules"
)

func main() {
	urlFlag := flag.String("url", "", "URL to scan for XSS (e.g. http://site.com/page?param=value)")
	formatFlag := flag.String("format", "json", "Report format: json, html, csv")
	output := flag.String("output", "report.json", "Report file name")
	proxyURL := flag.String("proxy", "", "Proxy URL (e.g. http://127.0.0.1:8080)")
	userAgent := flag.String("user-agent", "XSSScanner/1.0", "User-Agent string")
	maxWorkers := flag.Int("workers", 10, "Number of concurrent workers")
	maxPayloadsPerParam := flag.Int("max-payloads", 100, "Maximum number of payloads to test per parameter")
	enableScreenshots := flag.Bool("screenshots", false, "Enable screenshots")
	requestTimeout := flag.Duration("timeout", 10*time.Second, "Request timeout")
	rateLimit := flag.Int("rate-limit", 100, "Rate limit (requests per second)")
	rateLimitInterval := flag.Duration("rate-limit-interval", time.Second, "Rate limit interval")
	rateLimitMaxTokens := flag.Int("rate-limit-max-tokens", 100, "Rate limit max tokens")

	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("Usage: xssscanner --url=http://target.com --format=html --output=report.html")
		os.Exit(1)
	}

	config := &modules.ScannerConfig{
		MaxWorkers:           *maxWorkers,
		MaxPayloadsPerParam:  *maxPayloadsPerParam,
		EnableScreenshots:    *enableScreenshots,
		UserAgent:            *userAgent,
		ProxyURL:             *proxyURL,
		RequestTimeout:       *requestTimeout,
		RateLimit:            *rateLimit,
		RateLimitInterval:    *rateLimitInterval,
		RateLimitMaxTokens:   *rateLimitMaxTokens,
	}
	scanner := modules.NewXSSScanner(config)

	vulnChan := make(chan *common.Vulnerability, 1000)
	go scanner.Scan(context.TODO(), &common.Target{URL: *urlFlag}, vulnChan)

	var results []common.Vulnerability
	for v := range vulnChan {
		results = append(results, *v)
	}

	resultCollector := scanner.Results()

	switch strings.ToLower(*formatFlag) {
	case "json":
		reports.ExportJSON(*output, results, resultCollector.Screenshots())
	case "html":
		reports.ExportHTML(*output, results, resultCollector.Screenshots())
	case "csv":
		reports.ExportCSV(*output, results)
	default:
		fmt.Println("Invalid format. Use 'json', 'html', or 'csv'")
	}
	fmt.Printf("[✔] Report saved to %s (%d vulnerabilities)\n", *output, len(results))
}
