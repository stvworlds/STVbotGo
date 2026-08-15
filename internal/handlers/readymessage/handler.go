package readymessage

import (
	"STVbotGo/internal/logger"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (h *readyHandler) onReady(s *discordgo.Session, _ *discordgo.Ready) {
	if h.startupChannelID == "" {
		logger.Info("startupChannelID is not configured, skipping startup message")
		return
	}
	startupDuration := time.Since(h.startedAt).Round(time.Millisecond)

	message := fmt.Sprintf(
		"✅ **Бот успішно запущено**\n⏱ Час запуску: `%s`",
		startupDuration,
	)

	if _, err := s.ChannelMessageSend(h.startupChannelID, message); err != nil {
		logger.WithError(err, "failed to send startup message")
	}
}
