package modules

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

func (x *XSSScanner) ScanDeserialization(ctx context.Context, url string, payload []byte) {
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := x.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// In a real implementation, we would need to check for signs of successful
	// deserialization, such as a 500 error or a specific response body.
	log.Printf("Deserialization scan response: %s", body)
}
