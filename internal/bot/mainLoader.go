package bot

import (
	"github.com/bwmarrin/discordgo"
)

func NewSession(token string) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsAll

	commandLoader(session)
	eventLoader(session)

	return session, nil
}
