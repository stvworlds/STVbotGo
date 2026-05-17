package ticketsMessages

import "github.com/bwmarrin/discordgo"

func ptr[T any](value T) *T {
	return &value
}

func OpenMessageComponents() []discordgo.MessageComponent {
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
