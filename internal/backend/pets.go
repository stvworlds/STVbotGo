package backend

import (
	"context"
	"encoding/json"
	"fmt"
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

	var cats []CatResponse

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&cats); err != nil {
		return "", err
	}

	if len(cats) == 0 {
		return "", fmt.Errorf("cat not found")
	}

	return cats[0].Url, nil
}
