package app

import (
	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers"
	"log"

	"github.com/bwmarrin/discordgo"
)

func session(token string, cfg config.Env) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsAll

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		handlers.RegisterCommands(s, s.State.User.ID)
	})

	handlers.RegisterAll(session, cfg)

	log.Println("Bot is now running. Press CTRL-C to exit.")
	return session, nil
}
