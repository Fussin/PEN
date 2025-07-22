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

func scanPostForm(target string, param string, payload string) (string, error) {
	data := url.Values{}
	data.Set(param, payload)
	req, _ := http.NewRequest("POST", target, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "AdvancedXSSScanner/1.0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b), nil
}

func scanPostJSON(target, param, payload string) (string, error) {
	body := map[string]string{param: payload}
	jsonBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", target, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "AdvancedXSSScanner/1.0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b), nil
}

func mutatePayloads(payload string) []string {
	return []string{
		payload,
		url.QueryEscape(payload),
		strings.ReplaceAll(payload, "<", "%3C"),
		strings.ReplaceAll(payload, "alert", "a"+"l"+"ert"),
		html.EscapeString(payload),
		`<scr<script>ipt>alert(1)</scr<script>ipt>`,
		strings.ToUpper(payload),
		strings.ToLower(payload),
	}
}
