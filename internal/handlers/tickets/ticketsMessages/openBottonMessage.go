package ticketsMessages

import (
	"STVbotGo/internal/handlers/tickets/ticketsForms"

	"github.com/bwmarrin/discordgo"
)

var ticketTypeTexts = map[ticketsForms.TicketType]string{
	ticketsForms.TicketSupport:     "🛟 Ви обрали категорію **Підтримка**.\n> Натисніть кнопку нижче, щоб відкрити форму звернення до підтримки.",
	ticketsForms.TicketAdminReport: "🎰 Ви обрали категорію **Скарга на адміністрацію**.\n> Натисніть кнопку нижче, щоб описати ситуацію.",
	ticketsForms.TicketUserReport:  "☠️ Ви обрали категорію **Скарга на гравця**.\n> Натисніть кнопку нижче, щоб подати скаргу.",
	ticketsForms.TicketMessage:     "📝 Ви обрали категорію **Повідомлення**.\n> Натисніть кнопку нижче, щоб залишити повідомлення адміністрації.",
}

func OpenBottonComponents(mtype string) []discordgo.MessageComponent {
	ticketType := ticketsForms.TicketType(mtype)

	text, ok := ticketTypeTexts[ticketType]
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
							CustomID: ticketsForms.ButtonID(ticketType),
						},
					},
				},
			},
		},
	}
}
