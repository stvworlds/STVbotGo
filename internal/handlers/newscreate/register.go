package newscreate

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, cfg config.Env) {
	s.AddHandler(commandNewsCreate(cfg))
}

func RegisterCommand(s *discordgo.Session, applicationID string) error {
	_, err := s.ApplicationCommandCreate(applicationID, "", &discordgo.ApplicationCommand{
		Name:        createNewsCommandName,
		Description: "Create news post",
	})
	return err
}
