package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const logsChannelID = "1385289387557916785"

func OnMemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	if _, err := s.ChannelMessageSend(logsChannelID, m.Mention()+" приєднався до діскорду"); err != nil {
		log.Printf("failed to send startup message: %v", err)
	}

	channel, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		log.Printf("Error creating dm channel: %v\n", err)
	}

	_, err = s.ChannelMessageSend(channel.ID, "Вітаю")
	if err != nil {
		log.Printf("Error on sending DM: %v\n", err)
	}

}
