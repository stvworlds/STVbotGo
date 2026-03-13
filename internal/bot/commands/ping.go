package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

const pingCommandName = "ping"

func CommandPing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if i.ApplicationCommandData().Name != pingCommandName {
		return
	}

	latency := s.HeartbeatLatency()
	content := fmt.Sprintf("🏓 Поточний пінг бота: `%dms`", latency.Milliseconds())

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content},
	})
	if err != nil {
		log.Printf("failed to respond to /ping command: %v", err)
	}
}
