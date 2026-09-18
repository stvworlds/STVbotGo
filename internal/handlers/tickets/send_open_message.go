package tickets

import (
	"STVbotGo/internal/handlers/tickets/ticketsMessages"
	"STVbotGo/internal/logger"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	channel                                           = "1395637822110695476"
	messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
)

func sendOpenTicketMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author == nil {
		logger.Debug("Message author is nil, ignoring.")
		return
	}

	if s.State == nil || s.State.User == nil {
		logger.Debug("Bot state user is nil, ignoring.")
		return
	}

	if m.Author.Bot {
		logger.Debug("Message sent by bot, ignoring.")
		return
	}

	if m.Author.ID == s.State.User.ID {
		logger.Debug("Message sent by bot, ignoring.")
		return
	}

	if m.GuildID == "" {
		logger.Debug("Message sent in DM, ignoring.")
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

	if strings.ToLower(m.Content) != "!tickets_resend_open_message" {
		logger.Debug("Invalid command, ignoring.")
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "Open ticket message has been resent to the channel.")
	if err != nil {
		return
	}

	_, err = s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Components: ticketsMessages.OpenMessageComponents(),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		logger.WithError(err, "Failed to send open ticket message:")
	}
}
