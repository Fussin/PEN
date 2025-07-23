package fingerprinting

import "strings"

type DBMSFingerprint struct {
	Name    string
	Pattern string
}

type DBMSFingerprinter struct {
	signatures []DBMSFingerprint
}

func NewDBMSFingerprinter() *DBMSFingerprinter {
	return &DBMSFingerprinter{
		signatures: []DBMSFingerprint{
			{Name: "MySQL", Pattern: "You have an error in your SQL syntax"},
			{Name: "MySQL", Pattern: "Supplied argument is not a valid MySQL result resource"},
			{Name: "PostgreSQL", Pattern: "unterminated quoted string"},
			{Name: "PostgreSQL", Pattern: "invalid input syntax for type"},
			{Name: "Microsoft SQL Server", Pattern: "Unclosed quotation mark"},
			{Name: "Microsoft SQL Server", Pattern: "An unhandled exception occurred during the execution of the current web request."},
			{Name: "Oracle", Pattern: "ORA-"},
			{Name: "Oracle", Pattern: "Oracle Error"},
			{Name: "SQLite", Pattern: "SQLite3::SQLException"},
			{Name: "SQLite", Pattern: "near \".\": syntax error"},
		},
	}
}

func (fp *DBMSFingerprinter) Match(body string) (bool, string, string) {
	lbody := strings.ToLower(body)
	for _, sig := range fp.signatures {
		if strings.Contains(lbody, strings.ToLower(sig.Pattern)) {
			return true, sig.Name, sig.Pattern
		}
	}
	return false, "", ""
}
