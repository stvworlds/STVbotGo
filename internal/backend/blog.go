package backend

import (
	"STVbotGo/internal/logger"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Post struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Anons     string `json:"anons"`
	Author    string `json:"author"`
	Thumbnail string `json:"thumbnail"`
	FullText  string `json:"full_text"`
	Time      string `json:"time"`
}

type CreateNewsRequest struct {
	Title     string `json:"title"`
	Anons     string `json:"anons"`
	FullText  string `json:"fullText"`
	Author    string `json:"author"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type CreateNewsResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
	Data    struct {
		ID int `json:"id"`
	} `json:"data"`
}

func (c *Client) CreatePost(
	ctx context.Context,
	request CreateNewsRequest,
) (*Post, error) {
	u, err := url.Parse(c.baseURL + "/blog")
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	resp, err := c.Request(ctx, u, http.MethodPost, body)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.WithError(err, "close blog response body")
		}
	}()

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	responseBody := string(responseBodyBytes)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf(
			"api returned status %d: %s",
			resp.StatusCode,
			responseBody,
		)
	}

	var createResponse CreateNewsResponse
	if err := json.Unmarshal(responseBodyBytes, &createResponse); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &Post{
		Id: int64(createResponse.Data.ID),
	}, nil
}
