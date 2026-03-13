package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func NewSession(token string) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is now running. Press CTRL-C to exit.")
	})
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	return session, nil
}
