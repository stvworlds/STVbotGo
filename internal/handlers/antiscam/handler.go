package antiscam

import (
	"STVbotGo/internal/logger"
	"STVbotGo/internal/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (h *cfg) antiScam(s *discordgo.Session, m *discordgo.MessageCreate) {

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
		logger.WithError(err, "failed to send message to actions log channel")
	}

	invite, err := s.ChannelInviteCreate(m.ChannelID, discordgo.Invite{
		MaxAge:  0,
		MaxUses: 1,
		Unique:  true,
	})
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to create invite for user %s", m.Author.ID))
	}
	inviteURL := "https://discord.gg/" + invite.Code

	message, err := sendUserMessageDM(s, m, m.Author, inviteURL)
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to send DM to user %s", m.Author.ID))
	}

	if err := utils.UserBan(m.GuildID, m.Author.ID, s, deleteMessageSeconds, "Anti-scam triggered"); err != nil {
		banStatus = "Failed"

		if deleteMessageErr := s.ChannelMessageDelete(message.ChannelID, message.ID); deleteMessageErr != nil {
			logger.WithError(deleteMessageErr, fmt.Sprintf("failed to delete message from user %s", m.Author.ID))
		}
		messageStatus = "Skipped"

		if _, deleteInviteErr := s.InviteDelete(invite.Code); deleteInviteErr != nil {
			logger.WithError(deleteInviteErr, fmt.Sprintf("failed to delete invite %s", invite.Code))
		}
	} else {
		banStatus = "Success"
		messageStatus = "Sent"
	}

	sendAdminMessage(s, m, h.ActionsLogChannelID, banStatus, messageStatus)
}
