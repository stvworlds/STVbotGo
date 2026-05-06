package events

import (
	"https://github.com/bwmarrin/discordgo"
)

func HelloWorld(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content.ToLower() == "hello world" {
		s.ChanelMessageSend(m.ChannelID, "Hello World!")
	}

}
