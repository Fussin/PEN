package reports

import (
	"encoding/csv"
	"encoding/json"
	"html/template"
	"os"
	"time"

	"github.com/autonomouspen/reconnaissance/internal/scanner/modules"
)

func ExportJSON(filename string, vulns []modules.Vulnerability) error {
	data, err := json.MarshalIndent(vulns, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func ExportCSV(filename string, vulns []modules.Vulnerability) error {
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

func ExportHTML(filename string, vulns []modules.Vulnerability) error {
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
			</tr>
			{{range .Vulns}}
			<tr>
				<td>{{.Type}}</td>
				<td>{{.Severity}}</td>
				<td>{{.URL}}</td>
				<td>{{.Parameter}}</td>
				<td>{{.Payload}}</td>
				<td>{{.Evidence}}</td>
				<td>{{.Confidence}}</td>
				<td>{{.CWE}}</td>
				<td>{{.CVSS}}</td>
				<td>{{.Timestamp}}</td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>
	`

	t := template.Must(template.New("report").Parse(tmpl))
	data := map[string]interface{}{
		"Vulns": vulns,
		"Count": len(vulns),
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.Execute(file, data)
}
