package antiscam

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func adminMessageComponents(u *discordgo.User, banStatus string, messageStatus string, disabled bool) []discordgo.MessageComponent {

	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{
				discordgo.Section{
					Components: []discordgo.MessageComponent{
						discordgo.TextDisplay{
							Content: "💀" + u.Mention() + " був відмічений системою як спамер. \n🔪 Cтатус покарання: `" + banStatus + "`\n✍️ Статус повідомлення: `" + messageStatus + "`",
						},
					},
					Accessory: discordgo.Thumbnail{
						Media: discordgo.UnfurledMediaItem{
							URL: u.AvatarURL("256"),
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
					Content: fmt.Sprintf("📜 Айді користувача: \n```\n%s\n```", u.ID),
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Зняти покарання",
					Style:    discordgo.PrimaryButton,
					CustomID: "anti_scam_ban_remove." + u.ID + "." + banStatus + "." + messageStatus,
					Disabled: disabled,
				},
			},
		},
	}
}
