package backend

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type CatResponse struct {
	Url string `json:"url"`
}

func (c *Client) RandomCat(ctx context.Context) (string, error) {
	u, err := url.Parse(c.baseURL + "/v1/images/search")
	if err != nil {
		return "", err
	}

	resp, err := c.Request(ctx, u, http.MethodGet, nil)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	switch resp.StatusCode {
	case http.StatusOK:
		var cat CatResponse
		if err := json.NewDecoder(resp.Body).Decode(&CatResponse{}); err != nil {
			return cat.Url, err
		}
		return cat.Url, nil
	default:
		return "", nil
	}
}
