package bot

import (
	"github.com/bwmarrin/discordgo"
)

func NewSession(token string) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.AddHandler(onReady)
	session.AddHandler(onInteractionCreate)
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	return session, nil
}
