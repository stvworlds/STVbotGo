package sync

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

func onBottonClick(s *discordgo.Session) {

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
	}
	logger.Debug(content)
	return fmt.Errorf("%v", err)
}
