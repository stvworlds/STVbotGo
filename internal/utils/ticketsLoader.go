package utils

import (
	"STVbotGo/internal/config"
	"STVbotGo/internal/tickets"
	"STVbotGo/internal/tickets/ticketsInteraction"
	"log"

	"github.com/bwmarrin/discordgo"
)

func ticketsLoader(s *discordgo.Session, cfg config.Env) {

	s.AddHandler(tickets.SendOpenTicketMessage)
	s.AddHandler(ticketsInteraction.OpenTicketButton)
	s.AddHandler(ticketsInteraction.TicketSelectMenu)
	s.AddHandler(ticketsInteraction.HandleTicketFormSubmit)
	log.Println("Tickets are now loaded")
}
