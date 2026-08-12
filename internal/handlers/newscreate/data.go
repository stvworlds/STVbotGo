package newscreate

import (
	"github.com/bwmarrin/discordgo"
)

const (
	createNewsCommandName = "create-news"
	createNewsModalID     = "create-news-modal"
	blogAPIURL            = "https://api.stvworlds.win/blog"
	blogPageBaseURL       = "https://stvworlds.win/blog/pages/"

	messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
)

func ptr[T any](value T) *T {
	return &value
}
