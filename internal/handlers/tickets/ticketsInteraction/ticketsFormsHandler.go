package ticketsInteraction

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"STVbotGo/internal/handlers/tickets/ticketsForms"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func OpenTicketButton(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	data := i.MessageComponentData()
	if !strings.HasPrefix(data.CustomID, ticketsForms.TicketButtonPrefix) {
		return
	}

	ticketType := ticketsData.TicketType(strings.TrimPrefix(data.CustomID, ticketsForms.TicketButtonPrefix))

	if ticketsForms.IsKnownForm(ticketType) {
		ticketsForms.OpenTicketForm(s, i, ticketType)
		return
	}
}
