package bot

import "github.com/bwmarrin/discordgo"
import "STVbotGo/internal/bot/commands"

func commandLoader(s *discordgo.Session) {
	s.AddHandler(commands.CommandPing)
	s.AddHandler(commands.CommandConvert)
}
