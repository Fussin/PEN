package detection

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/autonomouspen/scanner/internal/common"
)

type BooleanBased struct{}

func (b *BooleanBased) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	var findings []common.Finding

	payloads := []struct {
		True  string
		False string
	}{
		{"' AND 1=1 --", "' AND 1=2 --"},
		{"\" AND 1=1 --", "\" AND 1=2 --"},
		{"' OR 1=1 --", "' OR 1=2 --"},
		{"\" OR 1=1 --", "\" OR 1=2 --"},
	}

	for _, payload := range payloads {
		originalResp, err := http.Get(target)
		if err != nil {
			continue
		}
		originalBody, _ := ioutil.ReadAll(originalResp.Body)

		trueResp, err := http.Get(target + payload.True)
		if err != nil {
			continue
		}
		trueBody, _ := ioutil.ReadAll(trueResp.Body)

		falseResp, err := http.Get(target + payload.False)
		if err != nil {
			continue
		}
		falseBody, _ := ioutil.ReadAll(falseResp.Body)

		if !strings.EqualFold(string(originalBody), string(trueBody)) && strings.EqualFold(string(trueBody), string(falseBody)) {
			findings = append(findings, common.Finding{
				Type:       "Boolean-based SQLi",
				Severity:   "High",
				URL:        target + payload.True,
				Evidence:   "Response differs between true and false payloads.",
				Confidence: "High",
				CWE:        "CWE-89",
				CVSS:       8.8,
			})
		}
	}

	return findings, nil
}
