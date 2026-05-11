package events

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

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
	content := strings.TrimSpace(strings.ToLower(m.Content))

	found := false
	for _, name := range catsName {
		if content == name {
			found = true
			break
		}
	}
	if !found {
		return
	}

	resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		log.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения тела:", err)
		return
	}

	var catResp []CatResponse
	if err := json.Unmarshal(body, &catResp); err != nil {
		log.Println("Ошибка парсинга JSON:", err)
		return
	}

	if len(catResp) == 0 {
		return
	}

	catUrl := catResp[0].Url
	_, err = s.ChannelMessageSendReply(m.ChannelID, catUrl, m.Reference())
	if err != nil {
		log.Println("Ошибка отправки сообщения:", err)
		return
	}
}
