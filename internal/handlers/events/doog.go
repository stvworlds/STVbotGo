package events

import "github.com/bwmarrin/discordgo"

func ShowDoog(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!Собака" {
		_, err := s.ChannelMessageSend(m.ChannelID, "https://giphy.com/explore/smile-dog")
		if err != nil {
			return
		}
	}

}
