package readymessage

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, cfg config.Env) {
	readyHandler := newReadyHandler(cfg.StartupChannelID)
	s.AddHandler(readyHandler.onReady)
}

type readyHandler struct {
	startupChannelID string
}

func newReadyHandler(startupChannelID string) *readyHandler {
	return &readyHandler{
		startupChannelID: startupChannelID,
	}
}
