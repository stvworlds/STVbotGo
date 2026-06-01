package utils

import (
	"log"

	"STVbotGo/internal/handlers/convert"
	"STVbotGo/internal/handlers/newscreate"
	"STVbotGo/internal/handlers/ping"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(s *discordgo.Session) {
	applicationID := s.State.User.ID

	if err := ping.RegisterCommand(s, applicationID); err != nil {
		log.Printf("failed to register /ping command: %v", err)
	}

	if err := convert.RegisterCommand(s, applicationID); err != nil {
		log.Printf("failed to register /convert command: %v", err)
	}

	if err := newscreate.RegisterCommand(s, applicationID); err != nil {
		log.Printf("failed to register /create-news command: %v", err)
	}

	log.Println("Commands registered")
}
