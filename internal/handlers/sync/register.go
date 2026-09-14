package sync

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, _ config.Env) {
	s.AddHandler(onBottonClick)
	s.AddHandler(onUserJoin)
}
