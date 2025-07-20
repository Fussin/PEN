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
	config         *ScannerConfig
	client         *http.Client
	rateLimiter    *RateLimiter
	payloadManager *PayloadManager
	validator      *XSSValidator
	results        *ResultCollector
}

type ScannerConfig struct {
	MaxWorkers           int
	MaxPayloadsPerParam  int
	EnableScreenshots    bool
	UserAgent            string
	ProxyURL             string
	RequestTimeout       time.Duration
}

func NewXSSScanner(config *ScannerConfig) *XSSScanner {
	return &XSSScanner{
		config:         config,
		client:         createHTTPClient(config),
		rateLimiter:    NewRateLimiter(100, time.Second),
		payloadManager: NewPayloadManager(),
		validator:      NewXSSValidator(),
		results:        NewResultCollector(),
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
	payloads := x.payloadManager.GetPayloadsForContext("html")
	baseQuery := parsedURL.Query()

	for i, payload := range payloads {
		if i > x.config.MaxPayloadsPerParam {
			break
		}

		query := baseQuery
		query.Set(param, payload.Value)
		targetURL := *parsedURL
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
type Payload struct {
	Value string
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
