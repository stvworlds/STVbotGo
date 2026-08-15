package dog

import (
	"STVbotGo/internal/logger"
	"encoding/json"

	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func showDog(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.WithError(err, "Error closing response body")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var dogResp response
	err = json.Unmarshal(body, &dogResp)
	if err != nil {
		logger.WithError(err, "Error unmarshaling JSON response")
		return
	}

	dogUrl := dogResp.Message
	logger.Debugf("dog image url: %s", dogUrl)

	_, err = s.ChannelMessageSend(m.ChannelID, dogUrl)
	if err != nil {
		return
	}
}
