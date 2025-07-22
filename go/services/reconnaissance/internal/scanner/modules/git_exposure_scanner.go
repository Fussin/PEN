package modules

import (
	"context"
	"log"
	"net/http"
	"net/url"
)

func (x *XSSScanner) ScanGitExposure(ctx context.Context, targetURL string) {
	u, err := url.Parse(targetURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return
	}

	gitPaths := []string{
		"/.git/config",
		"/.git/HEAD",
	}

	for _, path := range gitPaths {
		req, _ := http.NewRequestWithContext(ctx, "GET", u.Scheme+"://"+u.Host+path, nil)
		resp, err := x.client.Do(req)
		if err != nil {
			log.Printf("Error sending request: %v", err)
			continue
		}
		if resp.StatusCode == 200 {
			log.Printf("Exposed Git file found: %s", path)
		}
	}
}
