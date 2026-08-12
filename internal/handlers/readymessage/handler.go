package readymessage

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (h *readyHandler) onReady(s *discordgo.Session, _ *discordgo.Ready) {
	if h.startupChannelID == "" {
		log.Println("startupChannelID is not configured, skipping startup message")
		return
	}
	startupDuration := time.Since(h.startedAt).Round(time.Millisecond)

	message := fmt.Sprintf(
		"✅ **Бот успішно запущено**\n⏱ Час запуску: `%s`",
		startupDuration,
	)

	if _, err := s.ChannelMessageSend(h.startupChannelID, message); err != nil {
		log.Printf("failed to send startup message: %v", err)
	}
}
