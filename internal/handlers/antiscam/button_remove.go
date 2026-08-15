package antiscam

import (
	"STVbotGo/internal/logger"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func removeBotton(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	data := i.MessageComponentData()
	if !strings.HasPrefix(data.CustomID, "anti_scam_ban_remove.") {
		return
	}

	if i.Member == nil || i.Member.User == nil {
		respondEphemeral(s, i, "❌ Не вдалося перевірити ваші права.")
		return
	}

	permissions, err := s.UserChannelPermissions(i.Member.User.ID, i.ChannelID)
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to get permissions for user %s", i.Member.User.ID))
		respondEphemeral(s, i, "❌ Не вдалося перевірити ваші права.")
		return
	}

	if permissions&discordgo.PermissionBanMembers == 0 {
		respondEphemeral(s, i, "❌ У вас немає прав на розбан.")
		return
	}

	parts := strings.Split(data.CustomID, ".")
	if len(parts) != 4 {
		logger.Warnf("Invalid anti-scam CustomID: %q, parts count: %d", data.CustomID, len(parts))
		respondEphemeral(s, i, "❌ Некоректні дані кнопки.")
		return
	}

	userID := parts[1]
	banStatus := parts[2]
	messageStatus := parts[3]

	err = s.GuildBanDelete(i.GuildID, userID)
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to unban user %s", userID))
		respondEphemeral(s, i, "❌ Не вдалося зняти блокування з користувача.")
		return
	}

	user, err := s.User(userID)
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to fetch user %s", userID))
		respondEphemeral(s, i, "❌ Не вдалося отримати інформацію про користувача.")
		return
	}

	components := adminMessageComponents(user, banStatus, messageStatus, true)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Components: components,
			Flags:      messageFlagsIsComponentsV2,
		},
	})
	if err != nil {
		logger.WithError(err, "failed to update anti-scam message")
	}
	respondEphemeral(s, i, user.Mention()+" розблоковано.")
}

func respondEphemeral(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		logger.WithError(err, "failed to send interaction response")
	}
}
