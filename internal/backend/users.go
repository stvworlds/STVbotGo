package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type User struct {
	ID        int64  `json:"id"`
	DiscordID string `json:"username"`
}

func (c *Client) FindUserByUsername(
	ctx context.Context,
	username string,
) (*User, error) {

	u, err := url.Parse(c.baseURL + "/user/find/nickname")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("username", username)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("backend returned status: %d, req: %v, req: %v, header: %v", resp.StatusCode, req.Body, req.Header, req.URL)
	}

	var user User

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
