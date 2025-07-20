package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/autonomouspen/reconnaissance/internal/reports"
	"github.com/autonomouspen/reconnaissance/internal/scanner/modules"
)

func main() {
	urlFlag := flag.String("url", "", "URL to scan for XSS (e.g. http://site.com/page?param=value)")
	formatFlag := flag.String("format", "json", "Report format: json, html, csv")
	output := flag.String("output", "report.json", "Report file name")

	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("Usage: xssscanner --url=http://target.com --format=html --output=report.html")
		os.Exit(1)
	}

	config := &modules.ScannerConfig{}
	scanner := modules.NewXSSScanner(config)

	err := scanner.ScanURL(context.TODO(), *urlFlag)
	if err != nil {
		fmt.Println("[!] Scan error:", err)
		os.Exit(1)
	}

	results := scanner.Results().All()

	switch strings.ToLower(*formatFlag) {
	case "json":
		reports.ExportJSON(*output, results)
	case "html":
		reports.ExportHTML(*output, results)
	case "csv":
		reports.ExportCSV(*output, results)
	default:
		fmt.Println("Invalid format. Use 'json', 'html', or 'csv'")
	}
	fmt.Printf("[✔] Report saved to %s (%d vulnerabilities)\n", *output, len(results))
}
