package reporting

import (
	"encoding/csv"
	"encoding/json"
	"html/template"
	"os"
	"time"

	"github.com/autonomouspen/scanner/internal/common"
)

type Reporter struct{}

func NewReporter() *Reporter {
	return &Reporter{}
}

func (r *Reporter) Generate(findings []common.Finding, format string, outputFile string) error {
	switch format {
	case "json":
		return r.generateJSON(findings, outputFile)
	case "csv":
		return r.generateCSV(findings, outputFile)
	case "html":
		return r.generateHTML(findings, outputFile)
	}
	return nil
}

func (r *Reporter) generateJSON(findings []common.Finding, outputFile string) error {
	data, err := json.MarshalIndent(findings, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(outputFile, data, 0644)
}

func (r *Reporter) generateCSV(findings []common.Finding, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Type", "Severity", "URL", "Evidence", "Confidence", "CWE", "CVSS", "Timestamp", "Parameter", "Payload", "DBMS"}
	writer.Write(header)

	for _, v := range findings {
		writer.Write([]string{
			v.Type, v.Severity, v.URL, v.Evidence, v.Confidence, v.CWE,
			"8.8", v.Timestamp.Format(time.RFC3339), v.Parameter, v.Payload, v.DBMS,
		})
	}
	return nil
}

func (r *Reporter) generateHTML(findings []common.Finding, outputFile string) error {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Scan Report</title>
		<style>
			body { font-family: sans-serif; }
			table { border-collapse: collapse; width: 100%; }
			th, td { border: 1px solid #ddd; padding: 8px; }
			th { background-color: #f2f2f2; }
			img { max-width: 100%; }
		</style>
	</head>
	<body>
		<h2>Scan Report ({{.Count}} Findings)</h2>
		<table>
			<tr>
				<th>Type</th>
				<th>Severity</th>
				<th>URL</th>
				<th>Evidence</th>
				<th>Confidence</th>
				<th>CWE</th>
				<th>CVSS</th>
				<th>Timestamp</th>
				<th>Parameter</th>
				<th>Payload</th>
				<th>DBMS</th>
			</tr>
			{{range .Findings}}
			<tr>
				<td>{{.Type}}</td>
				<td>{{.Severity}}</td>
				<td>{{.URL}}</td>
				<td><pre>{{.Evidence}}</pre></td>
				<td>{{.Confidence}}</td>
				<td>{{.CWE}}</td>
				<td>{{.CVSS}}</td>
				<td>{{.Timestamp}}</td>
				<td>{{.Parameter}}</td>
				<td><pre>{{.Payload}}</pre></td>
				<td>{{.DBMS}}</td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>
	`

	t := template.Must(template.New("report").Parse(tmpl))

	data := map[string]interface{}{
		"Findings": findings,
		"Count":    len(findings),
	}
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.Execute(file, data)
}

func (f *common.Finding) MarshalCSV() ([]string, error) {
	return []string{
		f.Type,
		f.Severity,
		f.URL,
		f.Evidence,
		f.Confidence,
		f.CWE,
		"8.8",
		f.Timestamp.Format(time.RFC3339),
		f.Parameter,
		f.Payload,
		f.DBMS,
	}, nil
}
