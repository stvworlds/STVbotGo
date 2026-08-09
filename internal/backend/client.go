package backend

import (
	"context"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    baseURL,
		apiKey:     apiKey,
	}
}

func (c *Client) Request(
	ctx context.Context,
	u *url.URL,
	method string,
) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", c.apiKey)

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
