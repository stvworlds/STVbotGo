package events

import (
	"encoding/json"

	"io"
	"log"
	"net/http"

	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func init() {
	register(func(s *discordgo.Session, _ config.Env) {
		s.AddHandler(ShowDoog)
	})
}

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
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var dogResp DogResponse
	json.Unmarshal(body, &dogResp)

	dogUrl := dogResp.Message
	log.Printf(dogUrl)

	_, err = s.ChannelMessageSend(m.ChannelID, dogUrl)
	if err != nil {
		return
	}
}
