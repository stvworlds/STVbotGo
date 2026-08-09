package antiscam

import (
	"github.com/bwmarrin/discordgo"
)

func userDMessage(m *discordgo.MessageCreate, url string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{
				discordgo.Section{
					Components: []discordgo.MessageComponent{
						discordgo.TextDisplay{
							Content: "💔 Шановний " + m.Author.Mention() + "!\n Наша система виявила, що ваше повідомлення може бути спамом або рекламою. Будь ласка, переконайтеся, що ваше повідомлення відповідає нашим правилам.",
						},
					},
					Accessory: discordgo.Thumbnail{
						Media: discordgo.UnfurledMediaItem{
							URL: "https://p.novaskin.me/5055813792.png?class=thumbnail",
						},
						Description: ptr("Scam info"),
					},
				},
			},
		},
		discordgo.Container{
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: "Скористайтесь кнопкою нижче, щоб повернутися на сервер.\nАбо напишіть на нашу пошту <support@stvworlds.win>.\n-# Запрошення одноразове з перманентним терміном дії",
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Запрошення",
							Style: discordgo.LinkButton,
							URL:   url,
							Emoji: &discordgo.ComponentEmoji{
								Name: "📌",
							},
						},
					},
				},
			},
		},
	}
}
