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

type UnionBasedEngine struct {
	client *httpclient.Client
}

func NewUnionBasedEngine(client *httpclient.Client) *UnionBasedEngine {
	return &UnionBasedEngine{
		client: client,
	}
}

func (e *UnionBasedEngine) Name() string {
	return "Union-Based SQLi"
}

func (e *UnionBasedEngine) Detect(target string, param common.InjectionPoint, scanner interface{}, ctx context.Context) (*common.Finding, error) {
	injectionMarker := "UNIONSELECT"

	for i := 1; i < 20; i++ {
		var columns []string
		for j := 0; j < i; j++ {
			columns = append(columns, fmt.Sprintf("'%s'", injectionMarker))
		}
		payload := fmt.Sprintf("' UNION SELECT %s --", strings.Join(columns, ","))
		injected := param.WithInjection(payload)

		resp, body, err := e.client.SendVerboseRequest(injected)
		if err != nil {
			continue
		}

		if strings.Contains(body, injectionMarker) {
			return &common.Finding{
				URL:        target,
				Parameter:  param.Name,
				Payload:    payload,
				Type:       e.Name(),
				Evidence:   "Injected data was found in the response.",
				Severity:   "High",
				Confidence: "High",
				Timestamp:  time.Now(),
				Request:    fmt.Sprintf("%+v", resp.Request),
				Response:   scannerutil.TruncateBody(body),
			}, nil
		}
	}

	return nil, nil
}
