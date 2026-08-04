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
	if i.GuildID != "1335083670364618772" {
		return
	}
	permissions, err := s.UserChannelPermissions(i.User.ID, i.ChannelID)
	if err != nil {
		return
	}

	if permissions&discordgo.PermissionAdministrator == 0 {
		return
	}

	content := fmt.Sprintf("Bebra")

	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		content = fmt.Sprintf("Не достатньо аргументів.")
		return
	}

	userID := options[0].UserValue(s).ID
	content = fmt.Sprintf("Bebra %s", userID)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content},
	})
	if err != nil {
		log.Printf("failed to respond to /whitelist command: %v", err)
	}
}
