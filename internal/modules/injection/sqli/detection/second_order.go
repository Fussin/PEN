package detection

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/httpclient"
	"github.com/autonomouspen/scanner/internal/scannerutil"
)

type SecondOrderEngine struct {
	client *httpclient.Client
}

func NewSecondOrderEngine(client *httpclient.Client) *SecondOrderEngine {
	return &SecondOrderEngine{
		client: client,
	}
}

func (e *SecondOrderEngine) Name() string {
	return "Second-Order SQLi"
}

func (e *SecondOrderEngine) Detect(target string, param common.InjectionPoint, scanner interface{}, ctx context.Context) (*common.Finding, error) {
	// In a real implementation, we would need a way to specify the submission
	// and verification endpoints. For now, we'll just use the same endpoint
	// for both.
	submissionURL := target
	verificationURL := target
	payload := "' AND 1=1 --"

	// Submit the payload
	injected := param.WithInjection(payload)
	_, _, err := e.client.SendVerboseRequest(injected)
	if err != nil {
		return nil, err
	}

	// Visit the verification page
	resp, body, err := e.client.SendVerboseRequest(param.WithoutInjection())
	if err != nil {
		return nil, err
	}

	if strings.Contains(body, "some indicator of successful injection") {
		return &common.Finding{
			URL:        submissionURL,
			Parameter:  param.Name,
			Payload:    payload,
			Type:       e.Name(),
			Evidence:   "Payload was reflected on a different page.",
			Severity:   "High",
			Confidence: "High",
			Timestamp:  time.Now(),
			Request:    fmt.Sprintf("%+v", resp.Request),
			Response:   scannerutil.TruncateBody(body),
		}, nil
	}

	return nil, nil
}
