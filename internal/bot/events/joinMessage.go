package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const logsChannelID = "1385289387557916785"
const messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15

func OnMemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	components := []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: "👋 " + m.Mention() + " приєднався до Discord!",
				},
			},
		},
	}

	_, err := s.ChannelMessageSendComplex(logsChannelID, &discordgo.MessageSend{
		Components: components,
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Printf("failed to send join message: %v", err)
	}

	channel, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		log.Printf("Error creating dm channel: %v\n", err)
		return
	}

	dmComponents := []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: "Вітаю на сервері, " + m.User.Username + "!",
				},
			},
		},
	}

	_, err = s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
		Components: dmComponents,
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Printf("Error on sending DM: %v\n", err)
	}
}
