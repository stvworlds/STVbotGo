package utils

import (
	"log"

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
		log.Printf("Failed to ban `%v` for %v: %v", userId, reason, err)
		return err
	}

	return nil
}
