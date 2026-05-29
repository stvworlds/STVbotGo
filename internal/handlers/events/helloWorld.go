package events

import (
	"strings"

	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func init() {
	register(func(s *discordgo.Session, _ config.Env) {
		s.AddHandler(HelloWorld)
	})
}

func HelloWorld(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.ToLower(m.Content) == "hello world" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Hello World!")
		if err != nil {
			return
		}
	}

}
