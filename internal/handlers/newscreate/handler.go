package newscreate

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func commandNewsCreate(cfg config.Env) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			handleCreateNewsCommand(s, i)
		case discordgo.InteractionModalSubmit:
			handleCreateNewsModalSubmit(s, i, cfg)
		}
	}
}
