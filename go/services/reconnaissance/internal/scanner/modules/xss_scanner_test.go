package modules

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/autonomouspen/reconnaissance/internal/common"
)

func TestXSSScanner(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("param")
		w.Write([]byte("<html><body>" + param + "</body></html>"))
	}))
	defer server.Close()

	config := &ScannerConfig{
		MaxWorkers:           1,
		MaxPayloadsPerParam:  1,
		EnableScreenshots:    false,
		UserAgent:            "test",
		RequestTimeout:       10 * time.Second,
		RateLimit:            100,
		RateLimitInterval:    time.Second,
		RateLimitMaxTokens:   100,
	}
	scanner := NewXSSScanner(config)
	scanner.payloadManager.contextPayloads["html"] = []common.Payload{
		{Value: "<script>alert('XSS')</script>"},
	}

	vulnChan := make(chan *common.Vulnerability, 1)
	go scanner.Scan(context.TODO(), &common.Target{URL: server.URL + "?param=test"}, vulnChan)

	select {
	case result := <-vulnChan:
		if result.Type != "Reflected XSS" {
			t.Errorf("Expected vulnerability type 'Reflected XSS', got '%s'", result.Type)
		}
	case <-time.After(5 * time.Second):
		t.Error("Timed out waiting for result")
	}
}
