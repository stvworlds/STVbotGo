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

type createNewsRequest struct {
	Title     string `json:"title"`
	Anons     string `json:"anons"`
	FullText  string `json:"fullText"`
	Author    string `json:"author"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type createNewsResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
	Data    struct {
		ID int `json:"id"`
	} `json:"data"`
}

func ptr[T any](value T) *T {
	return &value
}
