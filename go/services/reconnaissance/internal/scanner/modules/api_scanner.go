package modules

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (x *XSSScanner) ScanRESTAPI(ctx context.Context, method, url, payload string, body map[string]interface{}) {
	// Recursively replace all string values with the payload
	var replace func(data interface{})
	replace = func(data interface{}) {
		switch v := data.(type) {
		case map[string]interface{}:
			for key, val := range v {
				switch v2 := val.(type) {
				case string:
					v[key] = payload
				case map[string]interface{}, []interface{}:
					replace(v2)
				}
			}
		case []interface{}:
			for i, val := range v {
				switch v2 := val.(type) {
				case string:
					v[i] = payload
				case map[string]interface{}, []interface{}:
					replace(v2)
				}
			}
		}
	}
	replace(body)

	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := x.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	if x.validator.Validate(payload, string(respBody)) {
		log.Printf("XSS found in API response: %s", respBody)
	}
}

func (x *XSSScanner) ScanGraphQLAPI(ctx context.Context, url, payload string) {
	// In a real implementation, we would need to introspect the GraphQL schema
	// to identify all possible queries and mutations. For now, we'll just send
	// a simple query.
	query := `{"query": "query { __typename }"}`
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString(query))
	req.Header.Set("Content-Type", "application/json")

	resp, err := x.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	if x.validator.Validate(payload, string(respBody)) {
		log.Printf("XSS found in GraphQL response: %s", respBody)
	}
}
