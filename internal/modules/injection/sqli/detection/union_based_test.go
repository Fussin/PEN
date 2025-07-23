package detection

import (
	"context"
	"net/http"
	"testing"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/httpclient"
)

func TestUnionBasedDetect(t *testing.T) {
	// Mock HTTP Client
	mockClient := new(MockUnionClient)

	engine := NewUnionBasedEngine(mockClient)

	ip := common.InjectionPoint{
		Name:   "id",
		Method: "GET",
		URL:    "http://test.com/get_user?id=1",
	}

	finding, err := engine.Detect("http://test.com", ip, nil, context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if finding == nil {
		t.Fatal("Expected finding, got nil")
	}
}

type MockUnionClient struct{}

func (m *MockUnionClient) SendVerboseRequest(ip common.InjectionPoint) (*http.Response, string, error) {
	if ip.URL == "http://test.com/get_user?id=1' UNION SELECT 'UNIONSELECT' --" {
		return &http.Response{}, "UNIONSELECT", nil
	}
	return &http.Response{}, "", nil
}
