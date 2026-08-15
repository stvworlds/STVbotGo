package ping

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func commandPing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if i.ApplicationCommandData().Name != pingCommandName {
		return
	}

	latency := s.HeartbeatLatency()
	content := fmt.Sprintf("Pong: `%dms`", latency.Milliseconds())

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content},
	})
	if err != nil {
		log.Printf("failed to respond to /ping command: %v", err)
	}
}
