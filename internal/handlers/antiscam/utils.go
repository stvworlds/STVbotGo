package antiscam

import (
	"STVbotGo/internal/logger"
	"fmt"

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
		logger.WithError(err, "failed to send admin message")
	}
}

func sendUserMessageDM(s *discordgo.Session, m *discordgo.MessageCreate, u *discordgo.User, url string) (*discordgo.Message, error) {
	channel, err := s.UserChannelCreate(u.ID)
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to create DM channel for user %s", u.ID))
		return nil, err
	}

	res, err := s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
		Components: userDMessage(m, url),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		logger.WithError(err, "error sending DM")
		return nil, err
	}

	return res, nil
}
