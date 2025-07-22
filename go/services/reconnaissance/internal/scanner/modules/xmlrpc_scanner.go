package modules

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

func (x *XSSScanner) ScanXMLRPC(ctx context.Context, url string, payload string) {
	xml := `
		<?xml version="1.0"?>
		<methodCall>
			<methodName>test.method</methodName>
			<params>
				<param>
					<value><string>` + payload + `</string></value>
				</param>
			</params>
		</methodCall>
	`

	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString(xml))
	req.Header.Set("Content-Type", "text/xml")

	resp, err := x.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if x.validator.Validate(payload, string(body)) {
		log.Printf("XSS found in XML-RPC response: %s", body)
	}
}
