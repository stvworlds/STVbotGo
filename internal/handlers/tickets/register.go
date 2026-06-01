package tickets

import (
	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/tickets/ticketsInteraction"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, cfg config.Env) {
	s.AddHandler(sendOpenTicketMessage)
	s.AddHandler(ticketsInteraction.OpenTicketButton)
	s.AddHandler(ticketsInteraction.TicketSelectMenu)
	s.AddHandler(ticketsInteraction.HandleTicketFormSubmit)
	log.Println("Tickets are now loaded")
}
