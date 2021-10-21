package gsd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	ApiKey string
}

type Query map[string]interface{}

const BaseApiUrl string = "https://app.sportdataapi.com/api/v1"

func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
	}
}

// getPath formats a path and query params in preparation for a request
func getPath(path string, query Query) string {
	p, _ := url.Parse(path)
	q := p.Query()

	for k, v := range query {
		q.Set(k, fmt.Sprintf("%v", v))
	}

	p.RawQuery = q.Encode()
	return p.RequestURI()
}

func getUrl(path string) string {
	return fmt.Sprintf("%s%s", BaseApiUrl, path)
}

// Make an HTTP request using the DefaultClient
func (c *Client) MakeRequest(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, getUrl(path), body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", c.ApiKey)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
