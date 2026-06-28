package ticketsMessages

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"

	"github.com/bwmarrin/discordgo"
)

var ticketTexts = map[ticketsData.TicketType]map[string]string{
	ticketsData.TicketSupport: {
		"text":     "🛟 Ви обрали категорію **Підтримка**.\n> Натисніть кнопку нижче, щоб відкрити форму звернення до підтримки.",
		"category": "123",
	},
	ticketsData.TicketAdminReport: {"text": "🎰 Ви обрали категорію **Скарга на адміністрацію**.\n> Натисніть кнопку нижче, щоб описати ситуацію."},
	ticketsData.TicketUserReport:  {"text": "☠️ Ви обрали категорію **Скарга на гравця**.\n> Натисніть кнопку нижче, щоб подати скаргу."},
	ticketsData.TicketMessage:     {"text": "📝 Ви обрали категорію **Повідомлення**.\n> Натисніть кнопку нижче, щоб залишити повідомлення адміністрації."},
}

func OpenTicketMessageComponents(data ticketsData.TicketChannelData) []discordgo.MessageComponent {
	topic := data.Fields["topic"]
	description := data.Fields["description"]

	if topic == "" {
		topic = "Без теми"
	}

	if description == "" {
		description = "Опис не вказано."
	}

	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{

				discordgo.TextDisplay{
					Content: topic + "\n" + description,
				},

				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						/*
							discordgo.Button{
								Label:    "Прийняти",
								Style:    discordgo.SuccessButton,
								CustomID: tickets.TicketAcceptButtonID,
							},
						*/
						discordgo.Button{
							Label:    "Закрити",
							Style:    discordgo.DangerButton,
							CustomID: ticketsData.TicketCloseButtonID,
						},
					},
				},
			},
		},
	}
}
