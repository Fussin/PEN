package detection

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/httpclient"
)

func TestTimeBasedDetect(t *testing.T) {
	// Mock HTTP Client
	mockClient := new(MockTimeClient)

	engine := NewTimeBasedEngine(mockClient)

	ip := common.InjectionPoint{
		Name:   "username",
		Method: "GET",
		URL:    "http://test.com/login?username=admin",
	}

	finding, err := engine.Detect("http://test.com", ip, nil, context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if finding == nil {
		t.Fatal("Expected finding, got nil")
	}
}

type MockTimeClient struct{}

func (m *MockTimeClient) SendVerboseRequest(ip common.InjectionPoint) (*http.Response, string, error) {
	if ip.URL == "http://test.com/login?username=admin' AND SLEEP(5) --" {
		time.Sleep(6 * time.Second)
	}
	return &http.Response{}, "", nil
}
