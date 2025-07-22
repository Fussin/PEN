package common

import (
	"context"
	"time"
)

type Scanner interface {
	Init(config *ScannerConfig)
	Scan(target string, ctx context.Context) ([]Finding, error)
	Name() string
}

type Finding struct {
	Type        string
	Severity    string
	URL         string
	Evidence    string
	Confidence  string
	CWE         string
	CVSS        float64
	Timestamp   time.Time
	Request     string
	Response    string
	Screenshot  string
}

type ScannerConfig struct {
	MaxWorkers         int
	RequestTimeout     time.Duration
	ProxyURL           string
	UserAgent          string
	EnableScreenshots  bool
	RateLimit          int
	RateLimitInterval  time.Duration
	RateLimitMaxTokens int
	LogFile            string
	PayloadsFile       string
	EnabledScanners    []string
}

type Target struct {
	URL string
}
