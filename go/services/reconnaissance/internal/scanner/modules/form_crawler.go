package modules

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

type FormInput struct {
	Action string
	Method string
	Fields []string
}

func (x *XSSScanner) CrawlForms(pageURL string) ([]FormInput, error) {
	resp, err := x.client.Get(pageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var forms []FormInput
	doc.Find("form").Each(func(i int, s *goquery.Selection) {
		action, _ := s.Attr("action")
		method, _ := s.Attr("method")

		var inputs []string
		s.Find("input").Each(func(j int, input *goquery.Selection) {
			if name, exists := input.Attr("name"); exists {
				inputs = append(inputs, name)
			}
		})

		formURL, _ := url.Parse(pageURL)
		actionURL, err := formURL.Parse(action)
		if err == nil {
			forms = append(forms, FormInput{
				Action: actionURL.String(),
				Method: method,
				Fields: inputs,
			})
		}
	})

	return forms, nil
}
