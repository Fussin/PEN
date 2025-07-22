package modules

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// ContextAnalyzer detects HTML/JS/CSS context of an injection point
type ContextAnalyzer struct{}

func NewContextAnalyzer() *ContextAnalyzer {
	return &ContextAnalyzer{}
}

// AnalyzeContext determines where the parameter is reflected
func (c *ContextAnalyzer) AnalyzeContext(body string, paramName string, payload string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return "html" // fallback to html
	}

	// Check for script tags
	if doc.Find("script").Length() > 0 {
		return "javascript"
	}

	// Check for style tags
	if doc.Find("style").Length() > 0 {
		return "css"
	}

	// Check for attributes
	if doc.Find(payload).Length() > 0 {
		return "attribute"
	}

	return "html"
}
