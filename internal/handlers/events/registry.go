package events

import (
	"log"

	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

type Registrar func(s *discordgo.Session, cfg config.Env)

var registrars []Registrar

func register(registrar Registrar) {
	registrars = append(registrars, registrar)
}

func RegisterHandlers(s *discordgo.Session, cfg config.Env) {
	for _, registrar := range registrars {
		registrar(s, cfg)
	}

	log.Println("Events are now loaded.")
}
