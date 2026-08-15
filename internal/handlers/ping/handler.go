package ping

import (
	"STVbotGo/internal/backend"
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

	var url = "https://discord.com/api/webhooks/1538275361752875038/2uVpfg-_agvG-H9UVChyFK0V9hKIrXgYy9sYX1fyuw4BpXaZ6WqxKLSqejmJcMFsBEnZ"

	message := backend.Webhook{
		Username: new("BotUser"),
		Content:  new("This is a test message"),
	}

	err := backend.SendWebhook(url, message)
	if err != nil {
		log.Printf("failed to send webhook: %v", err)
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content},
	})
	if err != nil {
		log.Printf("failed to respond to /ping command: %v", err)
	}
}
