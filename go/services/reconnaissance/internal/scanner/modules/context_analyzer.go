package modules

import (
	"strings"
)

// ContextAnalyzer detects HTML/JS/CSS context of an injection point
type ContextAnalyzer struct{}

func NewContextAnalyzer() *ContextAnalyzer {
	return &ContextAnalyzer{}
}

// AnalyzeContext determines where the parameter is reflected
func (c *ContextAnalyzer) AnalyzeContext(body string, paramName string, payload string) string {
	// Very naive logic to start with
	if strings.Contains(body, "<script>") {
		return "javascript"
	}
	if strings.Contains(body, "<style>") {
		return "css"
	}
	if strings.Contains(body, "<a href=") {
		return "attribute"
	}
	return "html"
}
