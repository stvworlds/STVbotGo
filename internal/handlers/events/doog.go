package events

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type DogResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func ShowDoog(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "!Собака" {
		return
	}

	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		return
	}
	defer resp.Body.Close() // ← сразу после проверки

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var dogResp DogResponse // ← объявляем переменную
	json.Unmarshal(body, &dogResp)

	dogUrl := dogResp.Message
	fmt.Println(dogUrl)

	_, err = s.ChannelMessageSend(m.ChannelID, dogUrl)
	if err != nil {
		return
	}
}
