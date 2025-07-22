package modules

import (
	"bytes"
	"encoding/json"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func scanPostForm(client *http.Client, target string, param string, payload string) (string, error) {
	data := url.Values{}
	data.Set(param, payload)
	req, _ := http.NewRequest("POST", target, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "AdvancedXSSScanner/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b), nil
}

func scanPostJSON(client *http.Client, target string, body map[string]interface{}, payload string) (string, error) {
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
	req, _ := http.NewRequest("POST", target, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "AdvancedXSSScanner/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b), nil
}

func mutatePayloads(payload string) []string {
	var mutations []string
	mutations = append(mutations, payload)
	mutations = append(mutations, url.QueryEscape(payload))
	mutations = append(mutations, html.EscapeString(payload))
	mutations = append(mutations, strings.ToUpper(payload))
	mutations = append(mutations, strings.ToLower(payload))
	mutations = append(mutations, strings.ReplaceAll(payload, "a", "à"))
	mutations = append(mutations, strings.ReplaceAll(payload, "e", "é"))
	mutations = append(mutations, strings.ReplaceAll(payload, "i", "í"))
	mutations = append(mutations, strings.ReplaceAll(payload, "o", "ó"))
	mutations = append(mutations, strings.ReplaceAll(payload, "u", "ú"))
	return mutations
}
