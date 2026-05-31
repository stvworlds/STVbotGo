package dog

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, _ config.Env) {
	s.AddHandler(showDog)
}

type response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
