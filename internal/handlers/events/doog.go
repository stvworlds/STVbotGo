package events

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func ShowDoog(s *discordgo.Session, m *discordgo.MessageCreate) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")

	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	type DogResponse struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	json.Unmarshal(body, &dogResp)

	dogUrl := dogResp.Message
	fmt.Println(dogUrl)

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!Собака" {
		_, err := s.ChannelMessageSend(m.ChannelID, dogUrl)
		if err != nil {
			return
		}
	}

}
