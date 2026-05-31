package joinmessage

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onMemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	sendWelcomeDM(s, m)
}

func sendWelcomeDM(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	channel, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		log.Printf("Error creating DM channel: %v\n", err)
		return
	}

	_, err = s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
		Components: buildWelcomeComponents(m.User),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Printf("Error sending welcome DM: %v\n", err)
	}
}
