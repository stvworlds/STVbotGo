package utils

import (
	"log"

	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func NewSession(token string, cfg config.Env) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsAll

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		RegisterCommands(s)
	})

	commandLoader(session)
	eventsLoader(session, cfg)

	log.Println("Bot is now running. Press CTRL-C to exit.")
	return session, nil
}
