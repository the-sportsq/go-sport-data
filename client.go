package gsd

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	ApiKey string
}

const BaseApiUrl string = "https://app.sportdataapi.com/api/v1"

func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
	}
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
