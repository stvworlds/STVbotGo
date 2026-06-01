package joinmessage

import (
	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

const (
	messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15

	rulesURL   = "https://stvworlds.win/docs/rules"
	newsURL    = "https://stvworlds.win/blog"
	supportURL = "https://discord.com/channels/1335083670364618772/1395637822110695476"
	reviewsURL = "https://discord.com/channels/1335083670364618772/1443261484103630901"
)

func ptr[T any](value T) *T {
	return &value
}

func Register(s *discordgo.Session, _ config.Env) {
	s.AddHandler(onMemberJoin)
}
