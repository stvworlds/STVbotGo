package events

import (
	"STVbotGo/internal/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	register(func(s *discordgo.Session, cfg config.Env) {
		AntiScamConfig := antiSpamHandler(cfg.ScamChannelID, cfg.ActionsLogChannelID)
		s.AddHandler(AntiScamConfig.antiSpam)
	})
}

type AntiScamConfig struct {
	ScamChannelID       string `json:"1510035959033499709"`
	ActionsLogChannelID string
}

func antiSpamHandler(scamChannelID string, actionsLogChannelID string) *AntiScamConfig {
	return &AntiScamConfig{
		ScamChannelID:       scamChannelID,
		ActionsLogChannelID: actionsLogChannelID,
	}
}

const (
	deleteMessageSeconds = 60 * 60
)

func (h *AntiScamConfig) antiSpam(s *discordgo.Session, m *discordgo.MessageCreate) {

	banStatus := "Ok"

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Message.ChannelID != h.ScamChannelID {
		return
	}

	if len(m.Attachments) < 3 {
		return
	}

	msgSend := &discordgo.MessageSend{
		Reference: m.Message.Forward(),
	}
	_, err := s.ChannelMessageSendComplex(h.ActionsLogChannelID, msgSend)
	if err != nil {
		log.Printf("Failed to send message to actions log channel: %v", err)
		return
	}

	if err := banUser(s, m); err != nil {
		banStatus = "Failed"
	}

	sendUserMessage(s, m, h.ActionsLogChannelID, banStatus)
}

func banUser(s *discordgo.Session, m *discordgo.MessageCreate) error {
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
		return err
	}

	return nil
}

func sendUserMessage(s *discordgo.Session, m *discordgo.MessageCreate, channel string, status string) {
	_, err := s.ChannelMessageSend(channel, "User has been banned for anti-scam measures."+status)
	if err != nil {
		log.Printf("Failed to send admin message: %v", err)
	}
}
