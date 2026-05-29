package utils

import (
	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/events"

	"github.com/bwmarrin/discordgo"
)

func eventsLoader(s *discordgo.Session, cfg config.Env) {
	events.RegisterHandlers(s, cfg)
}
