package newscreate

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func buildNewsMessageComponents(news createNewsRequest, newsID string, newsURL string) []discordgo.MessageComponent {
	components := []discordgo.MessageComponent{
		buildNewsHeaderSection(news),
		discordgo.Separator{
			Divider: new(true),
		},
		discordgo.TextDisplay{
			Content: buildNewsDescription(news),
		},
		buildNewsActionsRow(newsURL),
		discordgo.Separator{
			Divider: new(true),
		},
	}

	if news.Thumbnail != "" {
		components = append(components, buildNewsImage(news))
	}

	components = append(
		components,
		discordgo.TextDisplay{
			Content: buildNewsPostData(news, newsID),
		},
		discordgo.TextDisplay{
			Content: "-# STEVIOUS • Новини сервера",
		},
	)

	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: components,
		},
	}
}

func buildNewsHeaderSection(news createNewsRequest) discordgo.MessageComponent {
	return discordgo.TextDisplay{
		Content: fmt.Sprintf(
			"## 📰 %s",
			news.Title,
		),
	}
}

func buildNewsDescription(news createNewsRequest) string {
	return fmt.Sprintf(
		"📝 Анонс \n%s",
		news.Anons,
	)
}

func buildNewsImage(news createNewsRequest) discordgo.MessageComponent {
	return discordgo.MediaGallery{
		Items: []discordgo.MediaGalleryItem{
			{
				Media: discordgo.UnfurledMediaItem{
					URL: news.Thumbnail,
				},
				Description: ptr("News image"),
			},
		},
	}
}

func buildNewsPostData(news createNewsRequest, newsID string) string {
	return fmt.Sprintf(
		"-# 👤 Автор <@%s> • 🆔 ID новини `%s`",
		news.Author,
		newsID,
	)
}

func buildNewsActionsRow(newsURL string) discordgo.MessageComponent {
	return discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label: "Детальніше",
				Style: discordgo.LinkButton,
				URL:   newsURL,
				Emoji: &discordgo.ComponentEmoji{
					Name: "🔗",
				},
			},
			discordgo.Button{
				Label: "Усі новини",
				Style: discordgo.LinkButton,
				URL:   "https://stvworlds.win/blog",
				Emoji: &discordgo.ComponentEmoji{
					Name: "📰",
				},
			},
		},
	}
}
