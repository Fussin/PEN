package modules

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (x *XSSScanner) ScanStoredXSS(basePageURL string, formPayload Payload, verifyPage string) error {
	forms, err := x.CrawlForms(basePageURL)
	if err != nil {
		return err
	}

	for _, form := range forms {
		data := url.Values{}
		for _, field := range form.Fields {
			data.Set(field, formPayload.Value)
		}

		req, _ := http.NewRequest(form.Method, form.Action, strings.NewReader(data.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("User-Agent", x.config.UserAgent)

		resp, err := x.client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// After submission, visit verify page
		time.Sleep(2 * time.Second) // wait for processing
		resp2, err := x.client.Get(verifyPage)
		if err != nil {
			continue
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
	}

	return nil
}
