package newscreate

import (
	"STVbotGo/internal/logger"

	"github.com/bwmarrin/discordgo"
)

func handleCreateNewsCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != createNewsCommandName {
		return
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: createNewsModalID,
			Title:    "Створення новини",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "title",
							Label:       "Заголовок",
							Style:       discordgo.TextInputShort,
							Placeholder: "Введіть заголовок новини",
							Required:    true,
							MaxLength:   256,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "anons",
							Label:       "Короткий опис",
							Style:       discordgo.TextInputParagraph,
							Placeholder: "Введіть короткий опис",
							Required:    true,
							MaxLength:   1000,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "fullText",
							Label:       "Повний текст",
							Style:       discordgo.TextInputParagraph,
							Placeholder: "Введіть повний текст новини",
							Required:    true,
							MaxLength:   4000,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "thumbnail",
							Label:       "URL фото",
							Style:       discordgo.TextInputShort,
							Placeholder: "https://example.com/image.png",
							Required:    false,
							MaxLength:   1000,
						},
					},
				},
			},
		},
	})
	if err != nil {
		logger.WithError(err, "failed to open /create-news modal")
	}
}
