package common

import (
	"context"
	"net/url"
	"strings"
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

type InjectionPoint struct {
	Name        string
	Method      string // GET, POST
	Location    string // param, header, cookie, json
	URL         string
	Body        string
	Headers     map[string]string
	ContentType string
}

func (p InjectionPoint) WithInjection(value string) InjectionPoint {
	// Example: replace only the given parameter
	return InjectionPoint{
		Name:        p.Name,
		Method:      p.Method,
		URL:         injectParam(p.URL, p.Name, value),
		Body:        injectBody(p.Body, p.Name, value),
		Headers:     p.Headers, // inject here too, if needed
		ContentType: p.ContentType,
		Location:    p.Location,
	}
}

func (p InjectionPoint) WithoutInjection() InjectionPoint {
	return p.WithInjection("SAFE_TEST") // base value for diffing
}

func injectParam(targetURL, param, value string) string {
	u, err := url.Parse(targetURL)
	if err != nil {
		return targetURL
	}
	q := u.Query()
	q.Set(param, value)
	u.RawQuery = q.Encode()
	return u.String()
}

func injectBody(body, param, value string) string {
	// This is a simplified example. A real implementation would need to handle
	// different content types, such as JSON and XML.
	return strings.ReplaceAll(body, "="+param, "="+value)
}
