package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/autonomouspen/scanner/internal/common"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (c *Client) SendRequest(ip common.InjectionPoint) (*http.Response, error) {
	req, err := http.NewRequest(ip.Method, ip.URL, bytes.NewBufferString(ip.Body))
	if err != nil {
		return nil, err
	}
	for key, value := range ip.Headers {
		req.Header.Set(key, value)
	}
	return c.client.Do(req)
}

func (c *Client) SendVerboseRequest(ip common.InjectionPoint) (*http.Response, string, error) {
	resp, err := c.SendRequest(ip)
	if err != nil {
		return nil, "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, "", err
	}
	return resp, string(body), nil
}
