package whitelist

import (
	"STVbotGo/internal/backend"
	"STVbotGo/internal/config"
	"STVbotGo/internal/logger"
	"STVbotGo/internal/utils"
	"context"
	"fmt"
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

	content := "Bebra"

	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		content = "Не достатньо аргументів."
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := backend.NewClient(
		"https://api.stvworlds.win",
		cfg.ApiKey,
	)

	argument := options[0].UserValue(s)

	var user *backend.User
	user, err = client.FindUserByUsername(ctx, argument.Username)
	if err != nil {
		content = fmt.Sprintf("Не вдалося знайти користувача. `%v`", err)

	} else if utils.Contains(user.Roles, "PLAYER") {
		content = fmt.Sprintf("%v вже має роль гравця", argument.Mention())
	} else {
		_, err = client.AddRole(ctx, user.ID, "PLAYER")
		if err != nil {
			content = fmt.Sprintf(
				"Не вдалося додати роль гравця. `%v`",
				err,
			)
		} else {
			content = fmt.Sprintf(
				"%v доданий до гравців",
				argument.Mention(),
			)
		}
	}
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content},
	})
	if err != nil {
		logger.WithError(err, "failed to respond to /whitelist command")
	}
}
