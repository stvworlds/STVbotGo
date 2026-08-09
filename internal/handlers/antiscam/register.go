package antiscam

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

const (
	messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
	deleteMessageSeconds                              = 60 * 60
)

func ptr[T any](value T) *T {
	return &value
}

func Register(s *discordgo.Session, cfg config.Env) {
	antiScamConfig := antiSpamHandler(cfg.ScamChannelID, cfg.ActionsLogChannelID)
	s.AddHandler(antiScamConfig.antiScam)
	s.AddHandler(removeBotton)
}

type cfg struct {
	ScamChannelID       string `json:"1510035959033499709"`
	ActionsLogChannelID string
}

func antiSpamHandler(scamChannelID string, actionsLogChannelID string) *cfg {
	return &cfg{
		ScamChannelID:       scamChannelID,
		ActionsLogChannelID: actionsLogChannelID,
	}
}
