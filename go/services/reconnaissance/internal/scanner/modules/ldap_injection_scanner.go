package modules

import (
	"context"
	"log"
	"net/http"
	"net/url"
)

func (x *XSSScanner) ScanLDAPInjection(ctx context.Context, targetURL string) {
	u, err := url.Parse(targetURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return
	}

	payloads := []string{
		"*",
		"*)(uid=*)",
		"x' or 1=1 or 'x'='y",
	}

	for key := range u.Query() {
		originalValue := u.Query().Get(key)
		for _, payload := range payloads {
			q := u.Query()
			q.Set(key, payload)
			u.RawQuery = q.Encode()
			req, _ := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
			resp, err := x.client.Do(req)
			if err != nil {
				log.Printf("Error sending request: %v", err)
				continue
			}
			// In a real implementation, we would need to check for signs of
			// successful LDAP injection, such as a specific error message or
			// a change in the response body.
			log.Printf("LDAP injection scan response for payload '%s': %s", payload, resp.Status)
		}
		q := u.Query()
		q.Set(key, originalValue)
		u.RawQuery = q.Encode()
	}
}
