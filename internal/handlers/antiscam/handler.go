package antiscam

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *cfg) antiSpam(s *discordgo.Session, m *discordgo.MessageCreate) {

	banStatus := "Unknown"
	messageStatus := "Unknown"

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
	}

	invite, err := s.ChannelInviteCreate(m.ChannelID, discordgo.Invite{
		MaxAge:  0,
		MaxUses: 1,
		Unique:  true,
	})
	if err != nil {
		log.Printf("Failed to create invite for user %s: %v", m.Author.ID, err)
	}
	inviteURL := "https://discord.gg/" + invite.Code

	message, err := sendUserMessageDM(s, m, m.Author, inviteURL)
	if err != nil {
		log.Printf("Failed to send DM to user %s: %v", m.Author.ID, err)
	}

	if err := banUser(s, m); err != nil {
		banStatus = "Failed"

		if deleteMessageErr := s.ChannelMessageDelete(message.ChannelID, message.ID); deleteMessageErr != nil {
			log.Printf("Failed to delete message from user %s: %v", m.Author.ID, deleteMessageErr)
		}
		messageStatus = "Skipped"

		if _, deleteInviteErr := s.InviteDelete(invite.Code); deleteInviteErr != nil {
			log.Printf("Failed to delete invite %s: %v", invite.Code, deleteInviteErr)
		}
	} else {
		banStatus = "Success"
	}

	sendAdminMessage(s, m, h.ActionsLogChannelID, banStatus, messageStatus)
}
