package modules

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func (x *XSSScanner) ScanCRLFInjection(ctx context.Context, targetURL string) {
	u, err := url.Parse(targetURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return
	}

	payloads := []string{
		"%0d%0aSet-Cookie:crlf=true",
		"/%0d%0aSet-Cookie:crlf=true",
	}

	for _, payload := range payloads {
		// Test in path
		req, _ := http.NewRequestWithContext(ctx, "GET", u.Scheme+"://"+u.Host+u.Path+payload, nil)
		resp, err := x.client.Do(req)
		if err != nil {
			log.Printf("Error sending request: %v", err)
			continue
		}
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "crlf" && cookie.Value == "true" {
				fmt.Printf("CRLF injection found in path: %s\n", u.Path+payload)
			}
		}

		// Test in query parameter
		for key := range u.Query() {
			originalValue := u.Query().Get(key)
			q := u.Query()
			q.Set(key, originalValue+payload)
			u.RawQuery = q.Encode()
			req, _ = http.NewRequestWithContext(ctx, "GET", u.String(), nil)
			resp, err = x.client.Do(req)
			if err != nil {
				log.Printf("Error sending request: %v", err)
				continue
			}
			for _, cookie := range resp.Cookies() {
				if cookie.Name == "crlf" && cookie.Value == "true" {
					fmt.Printf("CRLF injection found in query parameter '%s': %s\n", key, originalValue+payload)
				}
			}
			q.Set(key, originalValue)
			u.RawQuery = q.Encode()
		}
	}
}
