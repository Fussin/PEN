package modules

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (x *XSSScanner) ScanStoredXSS(basePageURL string, formPayload Payload, verifyPage string) error {
	data := url.Values{}
	data.Set("msg", formPayload.Value) // Assuming `msg` is the form's name

	req, _ := http.NewRequest("POST", basePageURL, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", x.config.UserAgent)

	resp, err := x.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// After submission, visit verify page
	time.Sleep(2 * time.Second) // wait for processing
	resp2, err := x.client.Get(verifyPage)
	if err != nil {
		return err
	}
	defer resp2.Body.Close()

	body, _ := ioutil.ReadAll(resp2.Body)

	if ok, evidence := x.validator.Validate(formPayload.Value, string(body)); ok {
		x.results.Add(Vulnerability{
			Type:      "Stored XSS",
			URL:       verifyPage,
			Payload:   formPayload.Value,
			Evidence:  evidence,
			Severity:  "High",
			Timestamp: time.Now(),
		})
	}

	return nil
}
