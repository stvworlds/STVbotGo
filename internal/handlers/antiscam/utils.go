package antiscam

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func sendAdminMessage(s *discordgo.Session, m *discordgo.MessageCreate, channel string, banStatus string, messageStatus string) {

	bottonDisabled := false

	if banStatus == "Failed" {
		bottonDisabled = true
	}

	_, err := s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Components: adminMessageComponents(m.Author, banStatus, messageStatus, bottonDisabled),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Println("Failed to send admin message:", err)
	}
}

func sendUserMessageDM(s *discordgo.Session, m *discordgo.MessageCreate, u *discordgo.User, url string) (*discordgo.Message, error) {
	channel, err := s.UserChannelCreate(u.ID)
	if err != nil {
		log.Printf("Failed to create DM channel for user %s: %v", u.ID, err)
		return nil, err
	}

	res, err := s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
		Components: userDMessage(m, url),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Printf("Error sending DM: %v\n", err)
		return nil, err
	}

	return res, nil
}
