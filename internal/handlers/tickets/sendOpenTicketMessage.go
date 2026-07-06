package tickets

import (
	"STVbotGo/internal/handlers/tickets/ticketsMessages"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	channel                                           = "1395637822110695476"
	messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
)

func sendOpenTicketMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		log.Println("Message sent by bot, ignoring.")
		return
	}

	if m.GuildID == "" {
		log.Println("Message sent in DM, ignoring.")
		return
	}

	permissions, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		log.Println("Failed to get user permissions:", err)
		return
	}

	if permissions&discordgo.PermissionAdministrator == 0 {
		log.Println("User does not have administrator permissions, ignoring.")
		return
	}

	if strings.ToLower(m.Content) != "!tickets_resend_open_message" {
		log.Println("Invalid command, ignoring.")
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
		log.Println("Failed to send open ticket message:", err)
	}
}
