package utils

import (
	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/commands/newsCommands"
	"log"

	"github.com/bwmarrin/discordgo"
)
import "STVbotGo/internal/handlers/commands"

func commandLoader(s *discordgo.Session, cfg config.Env) {
	s.AddHandler(commands.CommandPing)
	s.AddHandler(commands.CommandConvert)
	s.AddHandler(newsCommands.CommandNewsCreate(cfg))

	log.Println("Commands are now loaded.")
}
