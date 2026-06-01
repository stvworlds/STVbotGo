package ping

import "github.com/bwmarrin/discordgo"

const pingCommandName = "ping"

func Register(s *discordgo.Session) {
	s.AddHandler(commandPing)
}

func RegisterCommand(s *discordgo.Session, applicationID string) error {
	_, err := s.ApplicationCommandCreate(applicationID, "", &discordgo.ApplicationCommand{
		Name:        pingCommandName,
		Description: "Show bot ping",
	})
	return err
}
