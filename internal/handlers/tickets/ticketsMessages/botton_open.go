package ticketsMessages

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"STVbotGo/internal/handlers/tickets/ticketsForms"

	"github.com/bwmarrin/discordgo"
)

func OpenBottonComponents(mtype string) []discordgo.MessageComponent {
	title := "Натисніть кнопку, щоб відкрити форму подачі заявки."

	if ticketsData.Config != nil {
		if cfg, ok := ticketsData.Config.Tickets[mtype]; ok && cfg.Title != "" {
			title = "Ви обрали категорію **" + cfg.Title + "**.\n> Натисніть кнопку нижче, щоб відкрити форму."
		}
	}

	ticketType := mtype

	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{

				discordgo.TextDisplay{
					Content: title,
				},
			},
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
	}
}
