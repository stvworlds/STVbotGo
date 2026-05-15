package utils

import (
	"STVbotGo/internal/config"
	"STVbotGo/internal/tickets"
	"log"

	"github.com/bwmarrin/discordgo"
)

func ticketsLoader(s *discordgo.Session, cfg config.Env) {

	s.AddHandler(tickets.SendOpenTicketMessage)
	s.AddHandler(tickets.OpenTicketButton)
	s.AddHandler(tickets.TicketSelectMenu)
	log.Println("Tickets are now loaded")
}
