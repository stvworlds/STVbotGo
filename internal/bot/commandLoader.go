package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)
import "STVbotGo/internal/bot/commands"

func commandLoader(s *discordgo.Session) {
	s.AddHandler(commands.CommandPing)
	s.AddHandler(commands.CommandConvert)

	log.Println("Commands are now loaded.")
}
