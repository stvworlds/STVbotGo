package utils

import (
	"log"

	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/events"

	"github.com/bwmarrin/discordgo"
)

func eventsLoader(s *discordgo.Session, cfg config.Env) {

	readyHandler := events.NewReadyHandler(cfg.StartupChannelID)
	s.AddHandler(readyHandler.OnReady)

	s.AddHandler(events.OnMemberJoin)
	s.AddHandler(events.OnIpMessage)
	s.AddHandler(events.HelloWorld)

	log.Println("Events are now loaded.")
}
