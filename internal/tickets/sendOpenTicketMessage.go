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

	if m.GuildID == "" {
		return
	}

	permissions, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		log.Println("Failed to get user permissions:", err)
		return
	}

	if permissions&discordgo.PermissionAdministrator == 0 {
		return
	}

	if strings.ToLower(m.Content) != "!tickets_resend_open_message" {
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "Open ticket message has been resent to the channel.")
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
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{

				discordgo.Section{
					Components: []discordgo.MessageComponent{
						discordgo.TextDisplay{
							Content: "## 🎫 STV Support Center\n" +
								"\n" +
								"Виникла проблема або потрібна допомога?\n" +
								"Ми поруч! Просто обери категорію тікету нижче.\n" +
								"-# Або напишіть на нашу електронну пошту <support@stvworlds.win>\n",
						},
					},
					Accessory: discordgo.Thumbnail{
						Media: discordgo.UnfurledMediaItem{
							URL: "https://w7.pngwing.com/pngs/345/37/png-transparent-computer-icons-others-miscellaneous-blue-logo.png",
						},
						Description: ptr("Support"),
					},
				},

				discordgo.Separator{
					Divider: ptr(true),
				},

				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "stv_tickets_select",
							Placeholder: "📩 Обери категорію в меню нижче",
							Options: []discordgo.SelectMenuOption{
								{
									Label:       "Підтримка",
									Description: "Отримай допомогу з питаннями, багами або доступом.",
									Value:       "support",
									Emoji: &discordgo.ComponentEmoji{
										Name: "🛟",
									},
								},
								{
									Label:       "Скарга на гравця",
									Description: "Повідом про порушення, конфлікти або підозрілу поведінку.",
									Value:       "admin-report",
									Emoji: &discordgo.ComponentEmoji{
										Name: "🎰",
									},
								},
								{
									Label:       "Скарга на адміністрацію",
									Description: "Повідомити про порушення, надмірність або неправильну поведінку адміністрації.",
									Value:       "user-report",
									Emoji: &discordgo.ComponentEmoji{
										Name: "☠️",
									},
								},
								{
									Label:       "Повідомлення",
									Description: "Залиште повідомлення для адміністрації.",
									Value:       "message",
									Emoji: &discordgo.ComponentEmoji{
										Name: "📝",
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
