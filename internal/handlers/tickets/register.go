package tickets

import (
	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"STVbotGo/internal/handlers/tickets/ticketsInteraction"
	"STVbotGo/internal/logger"
	"os"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, cfg config.Env) {
	ticketsCfg, err := ticketsData.LoadTicketsConfig(cfg.TicketsConfigPath)
	if err != nil {
		logger.Errorf("failed to load tickets config: %v", err)
		os.Exit(1)
	}
	ticketsData.Config = ticketsCfg
	logger.Infof("Tickets config loaded: %d ticket types", len(ticketsCfg.Tickets))

	s.AddHandler(sendOpenTicketMessage)
	s.AddHandler(ticketsInteraction.OpenTicketButton)
	s.AddHandler(ticketsInteraction.TicketSelectMenu)
	s.AddHandler(ticketsInteraction.HandleTicketFormSubmit)
	s.AddHandler(ticketsInteraction.TicketControlButtons(cfg))
	logger.Info("Tickets are now loaded")
}
