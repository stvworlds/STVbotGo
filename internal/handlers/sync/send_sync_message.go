package sync

import (
	"STVbotGo/internal/logger"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func sendSyncMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		logger.Debug("Message sent by bot, ignoring.")
		return
	}

	if m.GuildID == "" {
		logger.Debug("Message sent in DM, ignoring.")
		return
	}

	permissions, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		logger.WithError(err, "Failed to get user permissions:")
		return
	}

	if permissions&discordgo.PermissionAdministrator == 0 {
		logger.Debug("User does not have administrator permissions, ignoring.")
		return
	}

	if strings.ToLower(m.Content) != ",hresend_sync_message" {
		logger.Debug("Invalid command, ignoring.")
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "Sync message has been resent to the channel.")
	if err != nil {
		return
	}

	_, err = s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Components: syncComponents(),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		logger.WithError(err, "Failed to send sync message:")
	}
}
