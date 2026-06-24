package ticketsForms

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const (
	formID = "tickets_form"
)

func createForm(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: formID,
			Title:    "Заполните анкету",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "description",
							Label:       "Че случилось",
							Style:       discordgo.TextInputParagraph,
							Placeholder: "опиши ситуацыю",
							Required:    true,
							MaxLength:   1000,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "nickname",
							Label:       "Ваш никнейм",
							Style:       discordgo.TextInputShort,
							Placeholder: "веди ник",
							Required:    true,
							MaxLength:   30,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "url",
							Label:       "Url",
							Style:       discordgo.TextInputShort,
							Placeholder: "вставь url на изображение",
							Required:    false,
							MaxLength:   500,
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Printf("Ошибка: %v", err)
	}
}
