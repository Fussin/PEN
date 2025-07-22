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
		if strings.Contains(doc.Find("script").Text(), payload) {
			return "javascript"
		}
	}

	// Check for style tags
	if doc.Find("style").Length() > 0 {
		if strings.Contains(doc.Find("style").Text(), payload) {
			return "css"
		}
	}

	// Check for attributes
	var attrContext bool
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		for _, attr := range s.Nodes[0].Attr {
			if strings.Contains(attr.Val, payload) {
				attrContext = true
			}
		}
	})
	if attrContext {
		return "attribute"
	}

	// Check for URL contexts
	if doc.Find("a[href*=\""+payload+"\"]").Length() > 0 {
		return "url"
	}

	return "html"
}
