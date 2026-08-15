package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Webhook struct {
	Username        *string          `json:"username,omitempty"`
	AvatarUrl       *string          `json:"avatar_url,omitempty"`
	Content         *string          `json:"content,omitempty"`
	Embeds          *[]Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
}

type Embed struct {
	Title       *string    `json:"title,omitempty"`
	Url         *string    `json:"url,omitempty"`
	Description *string    `json:"description,omitempty"`
	Color       *string    `json:"color,omitempty"`
	Author      *Author    `json:"author,omitempty"`
	Fields      *[]Field   `json:"fields,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
}

type Author struct {
	Name    *string `json:"name,omitempty"`
	Url     *string `json:"url,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   *string `json:"name,omitempty"`
	Value  *string `json:"value,omitempty"`
	Inline *bool   `json:"inline,omitempty"`
}

type Thumbnail struct {
	Url *string `json:"url,omitempty"`
}

type Image struct {
	Url *string `json:"url,omitempty"`
}

type Footer struct {
	Text    *string `json:"text,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
}

type AllowedMentions struct {
	Parse *[]string `json:"parse,omitempty"`
	Users *[]string `json:"users,omitempty"`
	Roles *[]string `json:"roles,omitempty"`
}

func SendWebhook(url string, message Webhook) error {
	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(resp.Body)

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(string(responseBody))
	}

	return nil
}
