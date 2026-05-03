package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const startupChannelID = "1385289387557916785"

func OnReady(s *discordgo.Session, _ *discordgo.Ready) {
	if _, err := s.ChannelMessageSend(startupChannelID, "✅ Бот увімкнувся та готовий до роботи!"); err != nil {
		log.Printf("failed to send startup message: %v", err)
	}
}
