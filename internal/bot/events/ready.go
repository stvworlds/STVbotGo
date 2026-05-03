package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const startupChannelID = "1385289387557916785"

func OnReady(s *discordgo.Session, _ *discordgo.Ready) {
	log.Println("Bot is now running. Press CTRL-C to exit.")

	if _, err := s.ChannelMessageSend(startupChannelID, "✅ Бот увімкнувся та готовий до роботи!"); err != nil {
		log.Printf("failed to send startup message: %v", err)
	}

	if _, err := s.ApplicationCommandCreate(s.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Показати пінг бота",
	}); err != nil {
		log.Printf("failed to register /ping command: %v", err)
	}

	if _, err := s.ApplicationCommandCreate(s.State.User.ID, "", &discordgo.ApplicationCommand{
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
}
