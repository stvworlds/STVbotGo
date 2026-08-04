package whitelist

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func commandWhitelist(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if i.ApplicationCommandData().Name != commandName {
		return
	}

	content := fmt.Sprintf("Bebra")

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content},
	})
	if err != nil {
		log.Printf("failed to respond to /whitelist command: %v", err)
	}
}
