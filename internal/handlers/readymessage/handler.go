package readymessage

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *readyHandler) onReady(s *discordgo.Session, _ *discordgo.Ready) {
	if h.startupChannelID == "" {
		log.Println("startupChannelID is not configured, skipping startup message")
		return
	}

	if _, err := s.ChannelMessageSend(h.startupChannelID, "✅ Бот увімкнувся та готовий до роботи !"); err != nil {
		log.Printf("failed to send startup message: %v", err)
	}
}
