package plugins

import "strings"

type PassivePlugin interface {
	Name() string
	Execute(url string, body string) []Finding
}

type Finding struct {
	Type     string
	Severity string
	URL      string
	Evidence string
}

type PasswordInputDetector struct{}

func (p *PasswordInputDetector) Name() string {
	return "PasswordFieldFinder"
}

func (p *PasswordInputDetector) Execute(url, body string) []Finding {
	if strings.Contains(body, "type=\"password\"") {
		return []Finding{{
			Type:     "Sensitive Input",
			Severity: "Info",
			URL:      url,
			Evidence: "Found <input type=\"password\">",
		}}
	}
	return nil
}
