package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(s *discordgo.Session) {
	applicationID := s.State.User.ID

	if _, err := s.ApplicationCommandCreate(applicationID, "", &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Показати пінг бота",
	}); err != nil {
		log.Printf("failed to register /ping command: %v", err)
	}

	if _, err := s.ApplicationCommandCreate(applicationID, "", &discordgo.ApplicationCommand{
		Name:        "convert",
		Description: "Конвертувати файл у GIF",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionAttachment,
				Name:        "file",
				Description: "Файл для конвертації: png, jpeg, jpg, mp4 або інший формат ffmpeg",
				Required:    true,
			},
		},
	}); err != nil {
		log.Printf("failed to register /convert command: %v", err)
	}

	if _, err := s.ApplicationCommandCreate(applicationID, "", &discordgo.ApplicationCommand{
		Name:        "create-news",
		Description: "Створити новину",
	}); err != nil {
		log.Printf("failed to register /create-news command: %v", err)
	}

	log.Println("Commands registered")
}
