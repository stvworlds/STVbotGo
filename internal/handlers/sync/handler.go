package sync

import (
	"STVbotGo/internal/backend"
	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"STVbotGo/internal/logger"
	"STVbotGo/internal/utils"
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func onButtonClick(s *discordgo.Session, i *discordgo.InteractionCreate, cfg config.Env) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	data := i.MessageComponentData()

	switch data.CustomID {
	case ticketsData.TicketCloseButtonID:
	case ButtonSyncID:
		handleSyncButton(s, i, cfg)
	}

}

func onUserJoin(s *discordgo.Session, e *discordgo.GuildMemberAdd, cfg config.Env) {
	if e.GuildID != "1335083670364618772" {
		return
	}
	err := syncUser(s, e.Member.User, e.GuildID, cfg)
	if err != nil {
		return
	}
}

func syncUser(s *discordgo.Session, user *discordgo.User, guildID string, cfg config.Env) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	content := "Виникла не відома помилка"

	client := backend.NewClient(
		"https://api.stvworlds.win",
		cfg.ApiKey,
	)

	var player *backend.User
	player, err = client.FindUserByUsername(ctx, user.Username)
	if err != nil {
		content = fmt.Sprintf("Не вдалося знайти користувача. `%v`", err)

	} else if utils.Contains(player.Roles, "PLAYER") {
		err = s.GuildMemberRoleAdd(guildID, user.ID, "1402775011814539375")
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		content = fmt.Sprintf("%v вже має роль гравця", user.Mention())
	} else {
		err = fmt.Errorf("користувач не має ролі PLAYER на сайті")
		content = err.Error()
	}
	logger.Debug(content)
	return err
}

func handleSyncButton(s *discordgo.Session, i *discordgo.InteractionCreate, cfg config.Env) {
	if i.GuildID != "1335083670364618772" {
		return
	}
	content := "Unknown error"

	err := syncUser(s, i.User, i.GuildID, cfg)
	if err != nil {
		content = fmt.Sprintf("Не вдалося синхронізувати користувача. `%v`", err)
		respondMessage(s, i, content, true)
		return
	}

	content = "Успішна синхронізація"
	respondMessage(s, i, content, true)
}

func respondMessage(s *discordgo.Session, i *discordgo.InteractionCreate, content string, ephemeral bool) {
	flags := discordgo.MessageFlags(0)
	if ephemeral {
		flags = discordgo.MessageFlagsEphemeral
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   flags,
		},
	}); err != nil {
		logger.WithError(err, "failed to respond sync interaction")
	}
}
