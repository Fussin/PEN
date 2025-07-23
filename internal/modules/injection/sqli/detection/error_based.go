package detection

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/autonomouspen/scanner/internal/common"
)

type ErrorBased struct{}

func (e *ErrorBased) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	var findings []common.Finding

	payloads := []string{
		"'",
		"\"",
		"\\",
		"()",
		"[]",
		"{}",
		"%",
	}

	errorPatterns := map[string]string{
		"MySQL":      "You have an error in your SQL syntax",
		"PostgreSQL": "syntax error at or near",
		"Microsoft SQL Server": "Unclosed quotation mark after the character string",
		"Oracle":     "ORA-00921: unexpected end of SQL command",
		"SQLite":     "near \".\": syntax error",
	}

	for _, payload := range payloads {
		resp, err := http.Get(target + payload)
		if err != nil {
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)

		for dbms, pattern := range errorPatterns {
			if strings.Contains(string(body), pattern) {
				findings = append(findings, common.Finding{
					Type:       "Error-based SQLi",
					Severity:   "High",
					URL:        target + payload,
					Evidence:   pattern,
					Confidence: "High",
					CWE:        "CWE-89",
					CVSS:       8.8,
					DBMS:       dbms,
				})
			}
		}
	}

	return findings, nil
}
