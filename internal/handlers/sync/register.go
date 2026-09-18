package sync

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

const (
	ButtonSyncID                                      = "sync_botton"
	channel                                           = "1548851761127620748"
	messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
)

func Register(s *discordgo.Session, cfg config.Env) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		onButtonClick(s, i, cfg)
	})
	s.AddHandler(func(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
		onUserJoin(s, e, cfg)
	})
	s.AddHandler(sendSyncMessage)
}
