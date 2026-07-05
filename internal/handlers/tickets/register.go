package tickets

import (
	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"STVbotGo/internal/handlers/tickets/ticketsInteraction"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, cfg config.Env) {
	ticketsCfg, err := ticketsData.LoadTicketsConfig(cfg.TicketsConfigPath)
	if err != nil {
		log.Fatalf("failed to load tickets config: %v", err)
	}
	ticketsData.Config = ticketsCfg
	log.Printf("Tickets config loaded: %d ticket types", len(ticketsCfg.Tickets))

	s.AddHandler(sendOpenTicketMessage)
	s.AddHandler(ticketsInteraction.OpenTicketButton)
	s.AddHandler(ticketsInteraction.TicketSelectMenu)
	s.AddHandler(ticketsInteraction.HandleTicketFormSubmit)
	s.AddHandler(ticketsInteraction.TicketControlButtons(cfg))
	log.Println("Tickets are now loaded")
}
