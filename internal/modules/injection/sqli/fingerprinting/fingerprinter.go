package fingerprinting

import (
	"net/http"
	"strings"
)

type DBMSFingerprinter struct{}

func NewDBMSFingerprinter() *DBMSFingerprinter {
	return &DBMSFingerprinter{}
}

func (f *DBMSFingerprinter) Fingerprint(resp *http.Response, body string) string {
	// In a real implementation, we would use a more sophisticated method
	// to fingerprint the DBMS. For now, we'll just check for some common
	// error messages.
	if strings.Contains(body, "mysql") {
		return "MySQL"
	}
	if strings.Contains(body, "postgresql") {
		return "PostgreSQL"
	}
	if strings.Contains(body, "microsoft sql server") {
		return "Microsoft SQL Server"
	}
	if strings.Contains(body, "oracle") {
		return "Oracle"
	}
	if strings.Contains(body, "sqlite") {
		return "SQLite"
	}
	return "Unknown"
}
