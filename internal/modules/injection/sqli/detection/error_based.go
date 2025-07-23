package detection

import (
	"context"
	"fmt"
	"time"

	"github.com/autonomouspen/scanner/internal/common"
	"github.com/autonomouspen/scanner/internal/httpclient"
	"github.com/autonomouspen/scanner/internal/modules/injection/sqli/fingerprinting"
	"github.com/autonomouspen/scanner/internal/modules/injection/sqli/payloads"
	"github.com/autonomouspen/scanner/internal/scannerutil"
)

type ErrorBasedEngine struct {
	fingerprinter *fingerprinting.DBMSFingerprinter
	payloads      []payloads.Payload
	client        *httpclient.Client
}

func NewErrorBasedEngine(fp *fingerprinting.DBMSFingerprinter, pm *payloads.Manager, client *httpclient.Client) *ErrorBasedEngine {
	return &ErrorBasedEngine{
		fingerprinter: fp,
		payloads:      pm.GetPayloads(),
		client:        client,
	}
}

func (e *ErrorBasedEngine) Name() string {
	return "Error-Based SQLi"
}

// Injects classic breaking payloads and checks for error messages from the DBMS
func (e *ErrorBasedEngine) Detect(target string, param common.InjectionPoint, scanner interface{}, ctx context.Context) (*common.Finding, error) {
	originalResp, _, _ := e.client.SendVerboseRequest(param.WithoutInjection())

	for _, p := range e.payloads {
		injected := param.WithInjection(p.Value)

		resp, body, err := e.client.SendVerboseRequest(injected)
		if err != nil {
			continue
		}

		if matched, dbms, signature := e.fingerprinter.Match(body); matched {
			diff := scannerutil.StringDifference(originalResp, body)
			f := &common.Finding{
				URL:        target,
				Parameter:  param.Name,
				Payload:    p.Value,
				DBMS:       dbms,
				Type:       e.Name(),
				Evidence:   signature,
				// Diff:      diff,
				Severity:   "High",
				Confidence: "High",
				Timestamp:  time.Now(),
				Request:    fmt.Sprintf("%+v", resp.Request),
				Response:   scannerutil.TruncateBody(body),
			}
			return f, nil
		}
	}

	return nil, nil
}
