package modules

import (
	"net/http"
	"strings"
)

type WAFFingerprinter struct{}

func NewWAFFingerprinter() *WAFFingerprinter {
	return &WAFFingerprinter{}
}

func (w *WAFFingerprinter) DetectWAF(resp *http.Response, body string) bool {
	if resp.StatusCode == 403 || resp.StatusCode == 406 || resp.StatusCode == 501 {
		return true
	}
	knownWAFHeaders := []string{
		"X-Akamai-Session-Info", "X-Sucuri", "X-Distil-CS", "X-WAF-Detected",
		"X-Cloud-Trace-Context", "X-Cloud-Trace-Context", "X-AppGW-Application-Gateway",
		"X-Azure-Ref", "X-CDN-Loop", "X-Edge-IP", "X-EdgeConnect-MidMile-RTT",
	}

	for k := range resp.Header {
		for _, header := range knownWAFHeaders {
			if strings.Contains(k, header) {
				return true
			}
		}
	}

	wafPatterns := []string{
		"Access Denied", "Request blocked", "Security Firewall", "WAF", "Mod_Security",
		"You have been blocked", "403 Forbidden", "This request was blocked",
		"Cloudflare", "Incapsula", "Akamai", "Sucuri", "F5", "Barracuda",
	}

	for _, pattern := range wafPatterns {
		if strings.Contains(body, pattern) {
			return true
		}
	}

	return false
}
