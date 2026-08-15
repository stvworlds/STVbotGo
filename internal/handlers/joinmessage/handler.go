package joinmessage

import (
	"STVbotGo/internal/logger"

	"github.com/bwmarrin/discordgo"
)

func onMemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	sendWelcomeDM(s, m)
}

func sendWelcomeDM(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	channel, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		logger.WithError(err, "Error creating DM channel")
		return
	}

	_, err = s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
		Components: buildWelcomeComponents(m.User),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		logger.WithError(err, "Error sending welcome DM")
	}
}
