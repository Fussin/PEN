package detection

import (
	"context"
	"net/http"
	"testing"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/httpclient"
	"github.com/autonomouspen/scanner/internal/modules/injection/sqli/fingerprinting"
	"github.com/autonomouspen/scanner/internal/modules/injection/sqli/payloads"
)

func TestErrorBasedDetect(t *testing.T) {
	// Mock HTTP Client
	mockClient := new(MockClient)
	mockFingerprinter := fingerprinting.NewDBMSFingerprinter()

	pm := payloads.NewManager("payloads.json")

	engine := NewErrorBasedEngine(mockFingerprinter, pm, mockClient)

	ip := common.InjectionPoint{
		Name:   "username",
		Method: "GET",
		URL:    "http://testphp.vulnweb.com/login.php?username=admin",
	}

	finding, err := engine.Detect("http://test", ip, nil, context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if finding == nil {
		t.Fatal("Expected finding, got nil")
	}
	if finding.DBMS == "" {
		t.Error("Expected DBMS fingerprint, got empty")
	}
}

type MockClient struct{}

func (m *MockClient) SendRequest(ip common.InjectionPoint) (*http.Response, error) {
	return &http.Response{Body: http.NoBody}, nil
}
func (m *MockClient) SendVerboseRequest(ip common.InjectionPoint) (*http.Response, string, error) {
	return &http.Response{}, "You have an error in your SQL syntax", nil
}
