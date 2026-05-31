package antiscam

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *cfg) antiSpam(s *discordgo.Session, m *discordgo.MessageCreate) {

	banStatus := "Ok"

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Message.ChannelID != h.ScamChannelID {
		return
	}

	if len(m.Attachments) < 2 {
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

	sendAdminMessage(s, m, h.ActionsLogChannelID, banStatus)
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

func sendAdminMessage(s *discordgo.Session, m *discordgo.MessageCreate, channel string, status string) {
	_, err := s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Components: adminMessageComponents(m, status),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Println("Failed to send admin message:", err)
		return
	}
}
