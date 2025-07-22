package reports

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"html/template"
	"os"
	"time"

	"github.com/autonomouspen/reconnaissance/internal/common"
)

type ReportVulnerability struct {
	common.Vulnerability
	Screenshot string
}

func ExportJSON(filename string, vulns []common.Vulnerability, screenshots map[string][]byte) error {
	reportVulns := []ReportVulnerability{}
	for _, v := range vulns {
		rv := ReportVulnerability{
			Vulnerability: v,
		}
		if img, ok := screenshots[v.URL]; ok {
			rv.Screenshot = base64.StdEncoding.EncodeToString(img)
		}
		reportVulns = append(reportVulns, rv)
	}

	data, err := json.MarshalIndent(reportVulns, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func ExportCSV(filename string, vulns []common.Vulnerability) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Type", "Severity", "URL", "Param", "Payload", "Evidence", "Confidence", "CWE", "CVSS", "Timestamp"}
	writer.Write(header)

	for _, v := range vulns {
		writer.Write([]string{
			v.Type, v.Severity, v.URL, v.Parameter,
			v.Payload, v.Evidence, v.Confidence, v.CWE,
			"6.1", v.Timestamp.Format(time.RFC3339),
		})
	}
	return nil
}

func ExportHTML(filename string, vulns []common.Vulnerability, screenshots map[string][]byte) error {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>XSS Scan Report</title>
		<style>
			body { font-family: sans-serif; }
			table { border-collapse: collapse; width: 100%; }
			th, td { border: 1px solid #ddd; padding: 8px; }
			th { background-color: #f2f2f2; }
			img { max-width: 100%; }
		</style>
	</head>
	<body>
		<h2>XSS Scanner Report ({{.Count}} Findings)</h2>
		<table>
			<tr>
				<th>Type</th>
				<th>Severity</th>
				<th>URL</th>
				<th>Parameter</th>
				<th>Payload</th>
				<th>Evidence</th>
				<th>Confidence</th>
				<th>CWE</th>
				<th>CVSS</th>
				<th>Timestamp</th>
				<th>Screenshot</th>
			</tr>
			{{range .Vulns}}
			<tr>
				<td>{{.Vulnerability.Type}}</td>
				<td>{{.Vulnerability.Severity}}</td>
				<td>{{.Vulnerability.URL}}</td>
				<td>{{.Vulnerability.Parameter}}</td>
				<td><pre>{{.Vulnerability.Payload}}</pre></td>
				<td><pre>{{.Vulnerability.Evidence}}</pre></td>
				<td>{{.Vulnerability.Confidence}}</td>
				<td>{{.Vulnerability.CWE}}</td>
				<td>{{.Vulnerability.CVSS}}</td>
				<td>{{.Vulnerability.Timestamp}}</td>
				<td><img src="data:image/png;base64,{{.Screenshot}}"></td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>
	`

	t := template.Must(template.New("report").Parse(tmpl))

	reportVulns := []ReportVulnerability{}
	for _, v := range vulns {
		rv := ReportVulnerability{
			Vulnerability: v,
		}
		if img, ok := screenshots[v.URL]; ok {
			rv.Screenshot = base64.StdEncoding.EncodeToString(img)
		}
		reportVulns = append(reportVulns, rv)
	}

	data := map[string]interface{}{
		"Vulns": reportVulns,
		"Count": len(vulns),
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.Execute(file, data)
}
