package bot

import "github.com/bwmarrin/discordgo"
import "STVbotGo/internal/bot/events"

func eventLoader(s *discordgo.Session) {
	s.AddHandler(events.OnReady)
	s.AddHandler(events.OnMemberJoin)
	s.AddHandler(events.OnIpMessage)
}
