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
	ID         int      `json:"id"`
	DiscordID  string   `json:"discordID"`
	Username   string   `json:"username"`
	Email      string   `json:"email"`
	Roles      []string `json:"roles"`
	UUIDMC     string   `json:"uuidMC"`
	NicknameMC string   `json:"nicknameMC"`
	CreatedAt  string   `json:"createdAt"`
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

	switch resp.StatusCode {
	case http.StatusOK:
		var user User
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			return nil, err
		}
		return &user, nil
	case http.StatusNotFound:
		return nil, fmt.Errorf("user not found")
	default:
		return nil, fmt.Errorf("backend returned status: %d, req body: %v, req url: %v", resp.StatusCode, req.Body, req.URL)
	}
}
