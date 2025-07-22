package modules

import (
	"context"
	"log"
	"net/http"
	"net/url"
)

func (x *XSSScanner) ScanBackupFiles(ctx context.Context, targetURL string) {
	u, err := url.Parse(targetURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return
	}

	backupExtensions := []string{
		".bak",
		".old",
		".tmp",
		".swp",
		"~",
	}

	for _, ext := range backupExtensions {
		req, _ := http.NewRequestWithContext(ctx, "GET", u.Scheme+"://"+u.Host+u.Path+ext, nil)
		resp, err := x.client.Do(req)
		if err != nil {
			log.Printf("Error sending request: %v", err)
			continue
		}
		if resp.StatusCode == 200 {
			log.Printf("Backup file found: %s", u.Path+ext)
		}
	}
}
