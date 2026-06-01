package autothreadopen

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, cfg config.Env) {
	autoThreadOpenHandler := newAutoThreadOpenHandler(cfg.NewsChannels)
	s.AddHandler(autoThreadOpenHandler.autoThreadOpen)
}

type cfg struct {
	newsChannels map[string]struct{}
}

func newAutoThreadOpenHandler(newsChannels []string) *cfg {
	channels := make(map[string]struct{}, len(newsChannels))

	for _, channelID := range newsChannels {
		channels[channelID] = struct{}{}
	}

	return &cfg{
		newsChannels: channels,
	}
}
