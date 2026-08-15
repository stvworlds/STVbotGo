package handlers

import (
	"STVbotGo/internal/handlers/tickets"
	"STVbotGo/internal/handlers/whitelist"
	"STVbotGo/internal/logger"

	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/antiscam"
	"STVbotGo/internal/handlers/autothreadopen"
	"STVbotGo/internal/handlers/cats"
	"STVbotGo/internal/handlers/convert"
	"STVbotGo/internal/handlers/dog"
	"STVbotGo/internal/handlers/joinmessage"
	"STVbotGo/internal/handlers/newscreate"
	"STVbotGo/internal/handlers/ping"
	"STVbotGo/internal/handlers/readymessage"

	"github.com/bwmarrin/discordgo"
)

func RegisterAll(s *discordgo.Session, cfg config.Env) {
	ping.Register(s)
	convert.Register(s)
	newscreate.Register(s, cfg)

	antiscam.Register(s, cfg)
	autothreadopen.Register(s, cfg)
	cats.Register(s, cfg)
	dog.Register(s, cfg)
	joinmessage.Register(s, cfg)
	readymessage.Register(s, cfg)
	tickets.Register(s, cfg)

	whitelist.Register(s, cfg)
	logger.Info("Handlers are now loaded.")
}

func RegisterCommands(s *discordgo.Session, applicationID string) {
	if err := ping.RegisterCommand(s, applicationID); err != nil {
		logger.WithError(err, "failed to register /ping command")
	}
	if err := convert.RegisterCommand(s, applicationID); err != nil {
		logger.WithError(err, "failed to register /convert command")
	}
	if err := newscreate.RegisterCommand(s, applicationID); err != nil {
		logger.WithError(err, "failed to register /create-news command")
	}
	if err := whitelist.RegisterCommand(s, applicationID); err != nil {
		logger.WithError(err, "failed to register whitelist command")
	}
	logger.Info("Commands registered")
}
