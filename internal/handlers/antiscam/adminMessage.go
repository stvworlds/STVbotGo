package antiscam

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func adminMessageComponents(m *discordgo.MessageCreate, status string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{
				discordgo.Section{
					Components: []discordgo.MessageComponent{
						discordgo.TextDisplay{
							Content: m.Author.Mention() + " був відмічений системою як спамер. \nCтатус покарання: " + status,
						},
					},
					Accessory: discordgo.Thumbnail{
						Media: discordgo.UnfurledMediaItem{
							URL: m.Author.AvatarURL("256"),
						},
						Description: ptr("Scam Avatar"),
					},
				},
			},
		},
		discordgo.Container{
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: fmt.Sprintf("📜 Айді користувача: \n```\n%s\n```", m.Author.ID),
				},
			},
		},
		discordgo.Container{
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: fmt.Sprintf("🎰 Час приєднання до серверу:\n<t:%d:R>", m.Member.JoinedAt.Unix()),
				},
			},
		},
	}
}
