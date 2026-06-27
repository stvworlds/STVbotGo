package ticketsInteraction

import (
	"STVbotGo/internal/handlers/tickets/ticketsForms"
	"STVbotGo/internal/handlers/tickets/ticketsOpenLogic"

	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleTicketFormSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionModalSubmit {
		return
	}

	modalID := i.ModalSubmitData().CustomID
	if !isTicketModalID(modalID) {
		return
	}

	ticketType := extractTicketTypeFromModalID(modalID)
	if !ticketsForms.IsKnownForm(ticketType) {
		return
	}

	ticketsOpenLogic.CreateTickets(s, i, ticketsOpenLogic.TicketCreateData{
		Type:    ticketType,
		UserID:  interactionUserID(i),
		GuildID: i.GuildID,
		Fields:  getModalValues(i),
	})
}

func isTicketModalID(customID string) bool {
	return strings.HasPrefix(customID, ticketsForms.TicketModalPrefix) &&
		strings.HasSuffix(customID, ticketsForms.TicketModalSuffix)
}

func extractTicketTypeFromModalID(customID string) ticketsForms.TicketType {
	ticketType := strings.TrimPrefix(customID, ticketsForms.TicketModalPrefix)
	ticketType = strings.TrimSuffix(ticketType, ticketsForms.TicketModalSuffix)

	return ticketsForms.TicketType(ticketType)
}

func getModalValues(i *discordgo.InteractionCreate) map[string]string {
	values := make(map[string]string)

	for _, component := range i.ModalSubmitData().Components {
		row, ok := component.(*discordgo.ActionsRow)
		if !ok {
			continue
		}

		for _, rowComponent := range row.Components {
			input, ok := rowComponent.(*discordgo.TextInput)
			if !ok {
				continue
			}

			values[input.CustomID] = strings.TrimSpace(input.Value)
		}
	}

	return values
}

func interactionUserID(i *discordgo.InteractionCreate) string {
	if i.Member != nil && i.Member.User != nil {
		return i.Member.User.ID
	}

	if i.User != nil {
		return i.User.ID
	}

	return ""
}
