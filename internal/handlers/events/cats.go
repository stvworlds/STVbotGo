package events

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type CatResponse struct {
	Url string `json:"url"`
}

func ShowCat(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	catsName := []string{"котик", "котики", "кот", "коты", "кошка", "кошки", "кошечка", "кошечки"}

	found := false
	for _, name := range catsName {
		if m.Content == name {
			found = true
			break
		}
	}

	if !found {
		return
	}

	resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var catResp []CatResponse
	json.Unmarshal(body, &catResp)

	if len(catResp) == 0 {
		return
	}

	catUrl := catResp[0].Url // ✅ правильное поле
	fmt.Println(catUrl)

	_, err = s.ChannelMessageSend(m.ChannelID, catUrl)
	if err != nil {
		return
	}
}
