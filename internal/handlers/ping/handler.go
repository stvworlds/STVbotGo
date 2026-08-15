package ping

import (
	"STVbotGo/internal/logger"
	"fmt"

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
		logger.WithError(err, "failed to respond to /ping command")
	}
}
