package detection

import (
	"context"
	"fmt"
	"time"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/httpclient"
	"github.com/autonomouspen/scanner/internal/scannerutil"
)

type BooleanBasedEngine struct {
	client *httpclient.Client
}

func NewBooleanBasedEngine(client *httpclient.Client) *BooleanBasedEngine {
	return &BooleanBasedEngine{
		client: client,
	}
}

func (e *BooleanBasedEngine) Name() string {
	return "Boolean-Based SQLi"
}

func (e *BooleanBasedEngine) Detect(target string, param common.InjectionPoint, scanner interface{}, ctx context.Context) (*common.Finding, error) {
	payloads := []struct {
		True  string
		False string
	}{
		{"' AND 1=1 --", "' AND 1=2 --"},
		{"\" AND 1=1 --", "\" AND 1=2 --"},
		{"' OR 1=1 --", "' OR 1=2 --"},
		{"\" OR 1=1 --", "\" OR 1=2 --"},
	}

	originalResp, originalBody, err := e.client.SendVerboseRequest(param.WithoutInjection())
	if err != nil {
		return nil, err
	}

	for _, p := range payloads {
		trueInjected := param.WithInjection(p.True)
		_, trueBody, err := e.client.SendVerboseRequest(trueInjected)
		if err != nil {
			continue
		}

		falseInjected := param.WithInjection(p.False)
		_, falseBody, err := e.client.SendVerboseRequest(falseInjected)
		if err != nil {
			continue
		}

		if scannerutil.StringDifference(originalBody, trueBody) > 0.9 && scannerutil.StringDifference(trueBody, falseBody) > 0.9 {
			return &common.Finding{
				URL:        target,
				Parameter:  param.Name,
				Payload:    p.True,
				Type:       e.Name(),
				Evidence:   "Response differs between true and false payloads.",
				Severity:   "High",
				Confidence: "High",
				Timestamp:  time.Now(),
				Request:    fmt.Sprintf("%+v", originalResp.Request),
				Response:   scannerutil.TruncateBody(originalBody),
			}, nil
		}
	}

	return nil, nil
}
