package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)
import "STVbotGo/internal/handlers/commands"

func commandLoader(s *discordgo.Session) {
	s.AddHandler(commands.CommandPing)
	s.AddHandler(commands.CommandConvert)

	log.Println("Commands are now loaded.")
}
