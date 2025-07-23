package detection

import (
	"context"
	"net/http"
	"testing"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/httpclient"
)

func TestBooleanBasedDetect(t *testing.T) {
	// Mock HTTP Client
	mockClient := new(MockBooleanClient)

	engine := NewBooleanBasedEngine(mockClient)

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

type MockBooleanClient struct{}

func (m *MockBooleanClient) SendVerboseRequest(ip common.InjectionPoint) (*http.Response, string, error) {
	if ip.URL == "http://test.com/login?username=admin' AND 1=1 --" {
		return &http.Response{}, "Welcome admin", nil
	}
	if ip.URL == "http://test.com/login?username=admin' AND 1=2 --" {
		return &http.Response{}, "User not found", nil
	}
	return &http.Response{}, "Please login", nil
}
