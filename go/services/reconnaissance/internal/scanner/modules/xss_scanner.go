package modules

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/autonomouspen/reconnaissance/internal/common"
	"github.com/autonomouspen/reconnaissance/internal/scanner/plugins"
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
	passivePlugins  []plugins.PassivePlugin
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
	LogFile              string
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
		passivePlugins: []plugins.PassivePlugin{
			&plugins.PasswordInputDetector{},
		},
	}
}

type Task struct {
	URL   *url.URL
	Param string
}

func (x *XSSScanner) Scan(ctx context.Context, target *common.Target, vulnChan chan<- *common.Vulnerability) {
	parsedURL, err := url.Parse(target.URL)
	if err != nil {
		log.Printf("failed to parse URL: %v", err)
		return
	}

	tasks := make(chan Task, 100)
	var wg sync.WaitGroup

	for i := 0; i < x.config.MaxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range tasks {
				x.fuzzParameter(ctx, t.URL, t.Param, vulnChan)
			}
		}()
	}

	params := parsedURL.Query()
	for param := range params {
		tasks <- Task{URL: parsedURL, Param: param}
	}
	close(tasks)
	wg.Wait()
}

func (x *XSSScanner) fuzzParameter(ctx context.Context, parsedURL *url.URL, param string, vulnChan chan<- *common.Vulnerability) {
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
		log.Printf("Error creating request: %v", err)
		return
	}
	req.Header.Set("User-Agent", x.config.UserAgent)

	resp, err := x.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}
	body := string(bodyBytes)

	for _, plugin := range x.passivePlugins {
		for _, finding := range plugin.Execute(targetURL.String(), body) {
			vulnChan <- &common.Vulnerability{
				Type:     finding.Type,
				Severity: finding.Severity,
				URL:      finding.URL,
				Evidence: finding.Evidence,
			}
		}
	}

	if x.wafFingerprinter.DetectWAF(resp, body) {
		log.Println("[!] WAF Detected!")
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
			log.Printf("Error creating request: %v", err)
			continue
		}
		req.Header.Set("User-Agent", x.config.UserAgent)

		resp, err := x.client.Do(req)
		if err != nil {
			log.Printf("Error sending request: %v", err)
			continue
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			continue
		}

		body := string(bodyBytes)
		verified, evidence := x.validator.Validate(payload.Value, body)
		if verified {
			vuln := &common.Vulnerability{
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

			vulnChan <- vuln

			if x.config.EnableScreenshots {
				go func(v *common.Vulnerability) {
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
	proxyFunc := http.ProxyFromEnvironment
	if config.ProxyURL != "" {
		proxyURL, err := url.Parse(config.ProxyURL)
		if err == nil {
			proxyFunc = http.ProxyURL(proxyURL)
		}
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy:           proxyFunc,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: config.RequestTimeout,
	}
}

func (x *XSSScanner) Results() *ResultCollector {
	return x.results
}
