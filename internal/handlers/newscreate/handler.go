package newscreate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func commandNewsCreate(cfg config.Env) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			handleCreateNewsCommand(s, i)
		case discordgo.InteractionModalSubmit:
			handleCreateNewsModalSubmit(s, i, cfg)
		}
	}
}

func createNews(news createNewsRequest, apiKey string) (string, string, error) {
	body, err := json.Marshal(news)
	if err != nil {
		return "", "", fmt.Errorf("marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, blogAPIURL, bytes.NewReader(body))
	if err != nil {
		return "", "", fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", apiKey)

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("read response body: %w", err)
	}

	responseBody := string(responseBodyBytes)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return responseBody, "", fmt.Errorf("api returned status %d: %s", resp.StatusCode, responseBody)
	}

	newsID, err := extractNewsID(responseBodyBytes)
	if err != nil {
		return responseBody, "", err
	}

	return responseBody, newsID, nil
}
