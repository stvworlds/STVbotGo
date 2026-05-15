package tickets

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	channel                                           = "1501678737836736694"
	messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
)

func ptr[T any](value T) *T {
	return &value
}

func SendOpenTicketMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.ToLower(m.Content) != "!tickets_resend_open_message" {
		return
	}

	_, err := s.ChannelMessageSend(m.ChannelID, "Open ticket message has been resent to the channel.")
	if err != nil {
		return
	}

	_, err = s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Components: buildComponents(),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Println("Failed to send open ticket message:", err)
	}
}

func buildComponents() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{

				discordgo.TextDisplay{
					Content: "нажміте кнопку чтобі аткрить тікет",
				},

				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Аткрить",
							Style:    discordgo.SuccessButton,
							Disabled: false,
							CustomID: "stv_tickets_open",
						},
					},
				},

				discordgo.Separator{
					Divider: ptr(true),
				},

				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "stv_tickets_select",
							Placeholder: "Оберіть категорію",
							Options: []discordgo.SelectMenuOption{
								{
									Label: "Підтримка",
									Value: "pidtrimka",
									Emoji: &discordgo.ComponentEmoji{
										Name: "🦦",
									},
								},
								{
									Label: "Скарги",
									Value: "skargi",
									Emoji: &discordgo.ComponentEmoji{
										Name: "☠️",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
