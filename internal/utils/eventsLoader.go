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

	autoThreadOpenHandler := events.NewAutoThreadOpenHandler(cfg.NewsChannels)
	s.AddHandler(autoThreadOpenHandler.AutoThreadOpen)
	s.AddHandler(events.ShowDoog)
	log.Println("Events are now loaded.")
}
