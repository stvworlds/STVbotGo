package whitelist

import (
	"STVbotGo/internal/backend"
	"STVbotGo/internal/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func command(cfg config.Env) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commandHandler(s, i, cfg)
	}
}

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, cfg config.Env) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}
	if i.ApplicationCommandData().Name != commandName {
		return
	}
	if i.GuildID != "1335083670364618772" {
		return
	}
	if i.Member == nil {
		return
	}
	permissions, err := s.UserChannelPermissions(i.Member.User.ID, i.ChannelID)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	username := options[0].UserValue(s).Username

	client := backend.NewClient(
		"https://api.stvworlds.win",
		cfg.ApiKey,
	)

	var user *backend.User
	user, err = client.FindUserByUsername(ctx, username)
	if err != nil {
		log.Printf("failed to find user by username: %v", err)
		content = fmt.Sprintf("Не вдалося знайти користувача. %v", err)

	} else {
		content = fmt.Sprintf("Bebra %v", user.ID)
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content},
	})
	if err != nil {
		log.Printf("failed to respond to /whitelist command: %v", err)
	}
}
