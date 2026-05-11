package newsCommands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

const (
	createNewsCommandName = "create-news"
	createNewsModalID     = "create-news-modal"
	blogAPIURL            = "https://api.stvworlds.win/blog"
	blogPageBaseURL       = "https://stvworlds.win/blog/pages/"

	messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
)

type createNewsRequest struct {
	Title     string `json:"title"`
	Anons     string `json:"anons"`
	FullText  string `json:"fullText"`
	Author    string `json:"author"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type createNewsResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
	Data    struct {
		ID int `json:"id"`
	} `json:"data"`
}

func CommandNewsCreate(cfg config.Env) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			handleCreateNewsCommand(s, i)
		case discordgo.InteractionModalSubmit:
			handleCreateNewsModalSubmit(s, i, cfg)
		}
	}
}

func handleCreateNewsCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != createNewsCommandName {
		return
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: createNewsModalID,
			Title:    "Створення новини",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "title",
							Label:       "Заголовок",
							Style:       discordgo.TextInputShort,
							Placeholder: "Введіть заголовок новини",
							Required:    true,
							MaxLength:   256,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "anons",
							Label:       "Короткий опис",
							Style:       discordgo.TextInputParagraph,
							Placeholder: "Введіть короткий опис",
							Required:    true,
							MaxLength:   1000,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "fullText",
							Label:       "Повний текст",
							Style:       discordgo.TextInputParagraph,
							Placeholder: "Введіть повний текст новини",
							Required:    true,
							MaxLength:   4000,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "thumbnail",
							Label:       "URL фото",
							Style:       discordgo.TextInputShort,
							Placeholder: "https://example.com/image.png",
							Required:    false,
							MaxLength:   1000,
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Printf("failed to open /create-news modal: %v", err)
	}
}

func handleCreateNewsModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate, cfg config.Env) {
	if i.ModalSubmitData().CustomID != createNewsModalID {
		return
	}

	if cfg.BlogAPIKey == "" {
		respondEphemeral(s, i, "❌ API ключ для створення новин не налаштований.")
		return
	}

	if cfg.UserNewsChannel == "" {
		respondEphemeral(s, i, "❌ Канал для новин користувачів не налаштований.")
		return
	}

	authorID := getInteractionUserID(i)
	if authorID == "" {
		respondEphemeral(s, i, "❌ Не вдалося визначити автора новини.")
		return
	}

	news := createNewsRequest{
		Title:     sanitizeNewsInput(getModalValue(i, "title")),
		Anons:     sanitizeNewsInput(getModalValue(i, "anons")),
		FullText:  sanitizeNewsInput(getModalValue(i, "fullText")),
		Thumbnail: sanitizeNewsInput(getModalValue(i, "thumbnail")),
		Author:    authorID,
	}

	if news.Title == "" || news.Anons == "" || news.FullText == "" {
		respondEphemeral(s, i, "❌ Заповніть усі обовʼязкові поля.")
		return
	}

	if news.Thumbnail != "" && !isAllowedThumbnailURL(news.Thumbnail) {
		respondEphemeral(s, i, "❌ URL фото має починатися з `https` або `http`.")
		return
	}

	responseBody, newsID, err := createNews(news, cfg.BlogAPIKey)
	if err != nil {
		log.Printf("failed to create news: %v", err)
		respondEphemeral(s, i, fmt.Sprintf("❌ Не вдалося створити новину: %v", err))
		return
	}

	newsURL := blogPageBaseURL + newsID

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "✅ Новину успішно створено.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("failed to respond to create news modal: %v", err)
	}

	_, err = s.ChannelMessageSendComplex(cfg.UserNewsChannel, &discordgo.MessageSend{
		Components: buildNewsMessageComponents(news, newsID, newsURL),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Printf("failed to send news created component to channel %s: %v", cfg.UserNewsChannel, err)
		return
	}

	log.Printf("news created successfully, news id: %s, api response: %s", newsID, responseBody)
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

func extractNewsID(responseBody []byte) (string, error) {
	var response createNewsResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "", fmt.Errorf("parse create news response: %w", err)
	}

	if response.Status != "success" {
		return "", fmt.Errorf("api returned unsuccessful status: %s", response.Status)
	}

	if response.Data.ID <= 0 {
		return "", fmt.Errorf("api response does not contain valid news id")
	}

	return strconv.Itoa(response.Data.ID), nil
}

func buildNewsMessageComponents(news createNewsRequest, newsID string, newsURL string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{
				buildNewsHeaderSection(news),
				discordgo.Separator{
					Divider: new(true),
				},
				discordgo.TextDisplay{
					Content: buildNewsDescription(news, newsID),
				},
				discordgo.Separator{
					Divider: new(true),
				},
				buildNewsActionsRow(newsURL),
				discordgo.TextDisplay{
					Content: "-# STEVIOUS • Новини сервера",
				},
			},
		},
	}
}

func buildNewsHeaderSection(news createNewsRequest) discordgo.MessageComponent {
	if news.Thumbnail == "" {
		return discordgo.TextDisplay{
			Content: fmt.Sprintf(
				"## 📰 %s",
				news.Title,
			),
		}
	}

	return discordgo.Section{
		Components: []discordgo.MessageComponent{
			discordgo.TextDisplay{
				Content: fmt.Sprintf(
					"## 📰 %s",
					news.Title,
				),
			},
		},
		Accessory: discordgo.Thumbnail{
			Media: discordgo.UnfurledMediaItem{
				URL: news.Thumbnail,
			},
			Description: ptr("News thumbnail"),
		},
	}
}

func buildNewsDescription(news createNewsRequest, newsID string) string {
	return fmt.Sprintf(
		"**📝 %s\n\n"+
			"**👤 Автор** <@%s> • "+
			"**🆔 ID новини** `%s`",
		news.Anons,
		news.Author,
		newsID,
	)
}

func buildNewsActionsRow(newsURL string) discordgo.MessageComponent {
	return discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label: "Детальніше",
				Style: discordgo.LinkButton,
				URL:   newsURL,
				Emoji: &discordgo.ComponentEmoji{
					Name: "🔗",
				},
			},
			discordgo.Button{
				Label: "Усі новини",
				Style: discordgo.LinkButton,
				URL:   "https://stvworlds.win/blog",
				Emoji: &discordgo.ComponentEmoji{
					Name: "📰",
				},
			},
		},
	}
}

func sanitizeNewsInput(value string) string {
	value = strings.TrimSpace(value)

	replacer := strings.NewReplacer(
		`"`, "",
		`'`, "",
		`;`, "",
	)

	value = replacer.Replace(value)

	return strings.TrimSpace(value)
}

func isAllowedThumbnailURL(value string) bool {
	value = strings.ToLower(strings.TrimSpace(value))

	return strings.HasPrefix(value, "https://") || strings.HasPrefix(value, "http://")
}

func getInteractionUserID(i *discordgo.InteractionCreate) string {
	if i.Member != nil && i.Member.User != nil {
		return i.Member.User.ID
	}

	if i.User != nil {
		return i.User.ID
	}

	return ""
}

func getModalValue(i *discordgo.InteractionCreate, customID string) string {
	for _, component := range i.ModalSubmitData().Components {
		row, ok := component.(*discordgo.ActionsRow)
		if !ok {
			continue
		}

		for _, rowComponent := range row.Components {
			input, ok := rowComponent.(*discordgo.TextInput)
			if !ok {
				continue
			}

			if input.CustomID == customID {
				return strings.TrimSpace(input.Value)
			}
		}
	}

	return ""
}

func respondEphemeral(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("failed to send ephemeral response: %v", err)
	}
}

func ptr[T any](value T) *T {
	return &value
}
