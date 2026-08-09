package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type User struct {
	ID         int64    `json:"id"`
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

	resp, err := c.Request(ctx, u, http.MethodGet)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

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
		body, _ := io.ReadAll(resp.Body)
		apiError := fmt.Errorf(
			"backend returned status %d: %s",
			resp.StatusCode,
			strings.TrimSpace(string(body)))

		log.Println(apiError)
		return nil, apiError
	}
}

func (c *Client) AddRole(
	ctx context.Context,
	id int64,
	role string,
) (*User, error) {
	u, err := url.Parse(c.baseURL + "/user/roles")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("id", strconv.FormatInt(id, 10))
	q.Set("role", role)
	u.RawQuery = q.Encode()

	resp, err := c.Request(ctx, u, http.MethodPut)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

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
		body, _ := io.ReadAll(resp.Body)
		apiError := fmt.Errorf(
			"backend returned status %d: %s",
			resp.StatusCode,
			strings.TrimSpace(string(body)))

		log.Println(apiError)
		return nil, apiError
	}
}
