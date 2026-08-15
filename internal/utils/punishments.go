package utils

import (
	"STVbotGo/internal/logger"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func UserBan(guildId string, userId string, s *discordgo.Session, time int, reason string) error {
	data := struct {
		DeleteMessageSeconds int `json:"delete_message_seconds"`
	}{
		DeleteMessageSeconds: time,
	}

	_, err := s.RequestWithBucketID(
		"PUT",
		discordgo.EndpointGuildBan(guildId, userId),
		data,
		discordgo.EndpointGuildBan(userId, ""),
		discordgo.WithAuditLogReason(reason),
	)
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to ban `%v` for %v", userId, reason))
		return err
	}

	return nil
}
