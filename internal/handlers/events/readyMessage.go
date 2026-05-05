package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type ReadyHandler struct {
	StartupChannelID string
}

func NewReadyHandler(startupChannelID string) *ReadyHandler {
	return &ReadyHandler{
		StartupChannelID: startupChannelID,
	}
}

func (h *ReadyHandler) OnReady(s *discordgo.Session, _ *discordgo.Ready) {
	if h.StartupChannelID == "" {
		log.Println("startupChannelID is not configured, skipping startup message")
		return
	}

	if _, err := s.ChannelMessageSend(h.StartupChannelID, "✅ Бот увімкнувся та готовий до роботи"); err != nil {
		log.Printf("failed to send startup message: %v", err)
	}
}
