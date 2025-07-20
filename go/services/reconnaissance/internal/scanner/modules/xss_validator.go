package modules

import (
	"regexp"
	"strings"
)

// XSSValidator checks if a payload landed in an HTML context and executed
type XSSValidator struct {
	patterns []*regexp.Regexp
}

func NewXSSValidator() *XSSValidator {
	return &XSSValidator{
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)<script[^>]*?>[^<]*?alert\(`),
			regexp.MustCompile(`(?i)onerror\s*=\s*["']?alert\(`),
			regexp.MustCompile(`(?i)<img src=["']?x["']? onerror=`),
		},
	}
}

// Validate checks if the payload appears in the response
func (v *XSSValidator) Validate(payload, body string) (bool, string) {
	if strings.Contains(body, payload) {
		for _, pattern := range v.patterns {
			if pattern.MatchString(body) {
				return true, pattern.FindString(body)
			}
		}
	}
	return false, ""
}
