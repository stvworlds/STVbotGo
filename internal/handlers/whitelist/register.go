package whitelist

import "github.com/bwmarrin/discordgo"

const commandName = "whitelist"

func Register(s *discordgo.Session) {
	s.AddHandler(commandWhitelist)
}

func RegisterCommand(s *discordgo.Session, applicationID string) error {
	_, err := s.ApplicationCommandCreate(applicationID, "", &discordgo.ApplicationCommand{
		Name:        commandName,
		Description: "Add user to whitelist",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "Input user for adding to whitelist",
				Required:    true,
			},
		},
	})
	return err
}
