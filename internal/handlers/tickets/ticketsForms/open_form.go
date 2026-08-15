package ticketsForms

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"STVbotGo/internal/logger"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	TicketModalPrefix = "stv_tickets_"
	TicketModalSuffix = "_modal"

	TicketButtonPrefix = "stv_tickets_open_"
)

func ModalID(ticketType ticketsData.TicketType) string {
	return TicketModalPrefix + ticketType + TicketModalSuffix
}

func ButtonID(ticketType ticketsData.TicketType) string {
	return TicketButtonPrefix + ticketType
}

func IsKnownForm(ticketType ticketsData.TicketType) bool {
	if ticketsData.Config == nil {
		return false
	}
	_, ok := ticketsData.Config.Tickets[ticketType]
	return ok
}

func OpenTicketForm(s *discordgo.Session, i *discordgo.InteractionCreate, ticketType ticketsData.TicketType) {
	if ticketsData.Config == nil {
		logger.Info("tickets config not loaded")
		return
	}

	cfg, ok := ticketsData.Config.Resolve(ticketType)
	if !ok {
		logger.Infof("unknown ticket form type: %s", ticketType)
		return
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID:   ModalID(ticketType),
			Title:      cfg.Title,
			Components: ticketsData.BuildModalComponents(cfg.Modal.Fields),
		},
	})
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to open %s ticket modal", ticketType))
	}
}
