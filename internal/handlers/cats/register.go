package cats

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, _ config.Env) {
	s.AddHandler(showCat)
}

type CatResponse struct {
	Url string `json:"url"`
}
