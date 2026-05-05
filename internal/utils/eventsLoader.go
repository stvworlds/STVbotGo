package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

import (
	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/events"
)

func eventsLoader(s *discordgo.Session, cfg config.Env) {

	readyHandler := events.NewReadyHandler(cfg.StartupChannelID)
	s.AddHandler(readyHandler.OnReady)

	s.AddHandler(events.OnMemberJoin)
	s.AddHandler(events.OnIpMessage)

	log.Println("Events are now loaded.")
}
