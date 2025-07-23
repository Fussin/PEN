package detection

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/autonomouspen/scanner/internal/common"
)

type SecondOrder struct{}

func (s *SecondOrder) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	var findings []common.Finding
	// In a real implementation, we would need a way to specify the submission
	// and verification endpoints. For now, we'll just use the same endpoint
	// for both.
	submissionURL := target
	verificationURL := target
	payload := "' AND 1=1 --"

	// Submit the payload
	_, err := http.Get(submissionURL + payload)
	if err != nil {
		return nil, err
	}

	// Visit the verification page
	resp, err := http.Get(verificationURL)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(body), "some indicator of successful injection") {
		findings = append(findings, common.Finding{
			Type:       "Second-order SQLi",
			Severity:   "High",
			URL:        submissionURL + payload,
			Evidence:   "Payload was reflected on a different page.",
			Confidence: "High",
			CWE:        "CWE-89",
			CVSS:       8.8,
		})
	}

	return findings, nil
}
