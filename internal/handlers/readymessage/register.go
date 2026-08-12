package readymessage

import (
	"STVbotGo/internal/config"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, cfg config.Env) {
	readyHandler := newReadyHandler(cfg.StartupChannelID)
	s.AddHandler(readyHandler.onReady)
}

type readyHandler struct {
	startupChannelID string
	startedAt        time.Time
}

func newReadyHandler(startupChannelID string) *readyHandler {
	return &readyHandler{
		startupChannelID: startupChannelID,
		startedAt:        time.Now(),
	}
}
