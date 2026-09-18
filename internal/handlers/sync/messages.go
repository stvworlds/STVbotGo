package sync

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"

	"github.com/bwmarrin/discordgo"
)

func syncComponents() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: ticketsData.Ptr(6164128),
			Components: []discordgo.MessageComponent{

				discordgo.Section{
					Components: []discordgo.MessageComponent{
						discordgo.TextDisplay{
							Content: "## ⚾ Синхронізація доступу\n" +
								"\n" +
								"Натисніть на кнопку нище, щоб синхронізувати доступ (ролі) з нашим [сайтом](https://stvworlds.win)\n" +
								"-# За виникнення проблем, напишіть на нашу електронну пошту <support@stvworlds.win>, ало відкрийте заявку в https://discord.com/channels/1335083670364618772/1395637822110695476\n",
						},
					},
					Accessory: discordgo.Thumbnail{
						Media: discordgo.UnfurledMediaItem{
							URL: "https://th.bing.com/th/id/OIP.Gn8E4Ah_KM_CU_X-mV3MOwHaHa?r=0&o=7rm=3&rs=1&pid=ImgDetMain&o=7&rm=3",
						},
						Description: ticketsData.Ptr("Sync"),
					},
				},

				discordgo.Separator{
					Divider: ticketsData.Ptr(true),
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Синхронізувати ролі",
					Style:    discordgo.SuccessButton,
					Disabled: false,
					CustomID: ButtonSyncID,
				},
			},
		},
	}
}
