package ticketsInteraction

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var ticketTypeTexts = map[string]string{
	"stv_tickets_open_support":      "🛟 Ви обрали форму **Підтримка**.",
	"stv_tickets_open_admin-report": "🎰 Ви обрали форму **Скарга на гравця**.",
	"stv_tickets_open_user-report":  "☠️ Ви обрали форму **Скарга на адміністрацію**.",
	"stv_tickets_open_message":      "📝 Ви обрали форму **Повідомлення**.",
}

func message(text string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{

				discordgo.TextDisplay{
					Content: text + "\n я ебал потом долаем",
				},
			},
		},
	}
}

func OpenTicketButton(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	data := i.MessageComponentData()

	text, ok := ticketTypeTexts[data.CustomID]
	if !ok {
		text = "Натисніть кнопку, щоб відкрити форму подачі заявки."
	}

	switch data.CustomID {

	case "stv_tickets_open_support", "admin-report", "user-report", "message":
		{

			openTicket(message(text), s, i)

		}

	default:
		{
			return
		}
	}
}

func openTicket(message []discordgo.MessageComponent, s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Components: message,
			Flags:      messageFlagsIsComponentsV2 | discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("Error responding to interaction: %v", err)
		return
	}
}
