package modules

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type XSSScanner struct {
	config          *ScannerConfig
	client          *http.Client
	rateLimiter     *RateLimiter
	payloadManager  *PayloadManager
	contextAnalyzer *ContextAnalyzer
	validator       *XSSValidator
	results         *ResultCollector
	wafFingerprinter *WAFFingerprinter
}

type ScannerConfig struct {
	MaxWorkers           int
	MaxPayloadsPerParam  int
	EnableScreenshots    bool
	UserAgent            string
	ProxyURL             string
	RequestTimeout       time.Duration
	RateLimit            int
	RateLimitInterval    time.Duration
	RateLimitMaxTokens   int
}

func NewXSSScanner(config *ScannerConfig) *XSSScanner {
	pm := NewPayloadManager()
	pm.LoadPayloadsFromFile("payloads.json")
	return &XSSScanner{
		config:          config,
		client:          createHTTPClient(config),
		rateLimiter:     NewRateLimiter(config.RateLimit, config.RateLimitInterval, config.RateLimitMaxTokens),
		payloadManager:  pm,
		contextAnalyzer: NewContextAnalyzer(),
		validator:       NewXSSValidator(),
		results:         NewResultCollector(),
		wafFingerprinter: NewWAFFingerprinter(),
	}
}

type Task struct {
	URL   *url.URL
	Param string
}

func (x *XSSScanner) ScanURL(ctx context.Context, rawURL string) error {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %v", err)
	}

	tasks := make(chan Task, 100)
	var wg sync.WaitGroup

	for i := 0; i < x.config.MaxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range tasks {
				x.fuzzParameter(ctx, t.URL, t.Param)
			}
		}()
	}

	params := parsedURL.Query()
	for param := range params {
		tasks <- Task{URL: parsedURL, Param: param}
	}
	close(tasks)
	wg.Wait()
	return nil
}

func (x *XSSScanner) fuzzParameter(ctx context.Context, parsedURL *url.URL, param string) {
	baseQuery := parsedURL.Query()
	originalValue := baseQuery.Get(param)

	// First, send a request with a non-malicious payload to analyze the context
	testPayload := "test"
	query := baseQuery
	query.Set(param, testPayload)
	targetURL := *parsedURL
	targetURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", targetURL.String(), nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", x.config.UserAgent)

	resp, err := x.client.Do(req)
	if err != nil {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	body := string(bodyBytes)

	if x.wafFingerprinter.DetectWAF(resp, body) {
		fmt.Println("[!] WAF Detected!")
		// In a real implementation, we would adjust our scanning strategy here.
	}

	contextType := x.contextAnalyzer.AnalyzeContext(body, param, testPayload)
	payloads := x.payloadManager.GetPayloadsForContext(contextType)

	for i, payload := range payloads {
		if i > x.config.MaxPayloadsPerParam {
			break
		}

		query.Set(param, payload.Value)
		targetURL.RawQuery = query.Encode()

		req, err := http.NewRequestWithContext(ctx, "GET", targetURL.String(), nil)
		if err != nil {
			continue
		}
		req.Header.Set("User-Agent", x.config.UserAgent)

		resp, err := x.client.Do(req)
		if err != nil {
			continue
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			continue
		}

		body := string(bodyBytes)
		verified, evidence := x.validator.Validate(payload.Value, body)
		if verified {
			vuln := Vulnerability{
				Type:       "Reflected XSS",
				Severity:   "Medium",
				URL:        targetURL.String(),
				Parameter:  param,
				Payload:    payload.Value,
				Evidence:   evidence,
				Confidence: "High",
				CWE:        "CWE-79",
				CVSS:       6.1,
				Timestamp:  time.Now(),
				Request:    req.URL.String(),
				Response:   truncateBody(body),
			}

			x.results.Add(vuln)

			if x.config.EnableScreenshots {
				go func(v Vulnerability) {
					img := captureScreenshotWithChromedp(v.URL)
					x.results.AttachScreenshot(v.URL, img)
				}(vuln)
			}
		}
	}
	query.Set(param, originalValue)
	parsedURL.RawQuery = query.Encode()
}

func truncateBody(body string) string {
	max := 4096
	if len(body) > max {
		return body[:max] + "... [truncated]"
	}
	return body
}

func createHTTPClient(config *ScannerConfig) *http.Client {
	return &http.Client{}
}

func (x *XSSScanner) Results() *ResultCollector {
	return x.results
}
type Vulnerability struct {
	Type        string
	Severity    string
	URL         string
	Parameter   string
	Payload     string
	Evidence    string
	Confidence  string
	CWE         string
	CVSS        float64
	Timestamp   time.Time
	Request     string
	Response    string
}
