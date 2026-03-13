package bot

import "github.com/bwmarrin/discordgo"
import "STVbotGo/internal/bot/events"

func eventLoader(s *discordgo.Session) {
	s.AddHandler(events.OnReady)
}
