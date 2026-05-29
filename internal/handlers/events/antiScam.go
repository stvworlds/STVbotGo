package events

import (
	"STVbotGo/internal/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	register(func(s *discordgo.Session, _ config.Env) {
		s.AddHandler(antiSpam)
	})
}

func antiSpam(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Message.ChannelID != "1510035959033499709" {
		return
	}

	if len(m.Attachments) < 3 {
		return
	}

	deleteMessageSeconds := 60 * 60

	data := struct {
		DeleteMessageSeconds int `json:"delete_message_seconds"`
	}{
		DeleteMessageSeconds: deleteMessageSeconds,
	}

	_, err := s.RequestWithBucketID(
		"PUT",
		discordgo.EndpointGuildBan(m.GuildID, m.Author.ID),
		data,
		discordgo.EndpointGuildBan(m.GuildID, ""),
		discordgo.WithAuditLogReason("Anti-scam triggered"),
	)
	if err != nil {
		log.Printf("Failed to ban user for anti-scam: %v", err)
		return
	}
}
