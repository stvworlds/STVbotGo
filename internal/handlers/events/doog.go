package events

import (
	"encoding/json"
<<<<<<< HEAD

=======
>>>>>>> 04ab3f7ae38132b11c1eaa2f4afd035b77ae0956
	"io"
	"log"
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
