package events

import "github.com/bwmarrin/discordgo"

func ShowDoog(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!Собака" {
		_, err := s.ChannelMessageSend(m.ChannelID, "https://tenor.com/view/spearhyunho-gif-17972652055360680845")
		if err != nil {
			return
		}
	}

}
