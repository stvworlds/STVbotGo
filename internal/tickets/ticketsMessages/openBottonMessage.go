package ticketsMessages

import "github.com/bwmarrin/discordgo"

var ticketTypeTexts = map[string]string{
	"support":      "🛟 Ви обрали категорію **Підтримка**.\n> Натисніть кнопку нижче, щоб відкрити форму звернення до підтримки.",
	"admin-report": "🎰 Ви обрали категорію **Скарга на гравця**.\n> Натисніть кнопку нижче, щоб подати скаргу.",
	"user-report":  "☠️ Ви обрали категорію **Скарга на адміністрацію**.\n> Натисніть кнопку нижче, щоб описати ситуацію.",
	"message":      "📝 Ви обрали категорію **Повідомлення**.\n> Натисніть кнопку нижче, щоб залишити повідомлення адміністрації.",
}

func OpenBottonComponents(mtype string) []discordgo.MessageComponent {
	text, ok := ticketTypeTexts[mtype]
	if !ok {
		text = "Натисніть кнопку, щоб відкрити форму подачі заявки."
	}

	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{

				discordgo.TextDisplay{
					Content: text,
				},

				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Відкрити",
							Style:    discordgo.SuccessButton,
							Disabled: false,
							CustomID: "stv_tickets_open_" + mtype,
						},
					},
				},
			},
		},
	}
}
