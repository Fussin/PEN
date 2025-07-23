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
	Type        string    `json:"type"`
	Severity    string    `json:"severity"`
	URL         string    `json:"url"`
	Evidence    string    `json:"evidence"`
	Confidence  string    `json:"confidence"`
	CWE         string    `json:"cwe"`
	CVSS        float64   `json:"cvss"`
	Timestamp   time.Time `json:"timestamp"`
	Parameter   string    `json:"parameter,omitempty"`
	Payload     string    `json:"payload,omitempty"`
	DBMS        string    `json:"dbms,omitempty"`
	Screenshot  string    `json:"screenshot,omitempty"`
	Request     string    `json:"request,omitempty"`
	Response    string    `json:"response,omitempty"`
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
