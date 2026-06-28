package ticketsOpenLogic

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"STVbotGo/internal/handlers/tickets/ticketsMessages"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CreateTicketChannel(s *discordgo.Session, data ticketsData.TicketChannelData) (*discordgo.Channel, error) {
	if data.GuildID == "" {
		return nil, fmt.Errorf("guild ID is empty")
	}

	if data.UserID == "" {
		return nil, fmt.Errorf("ticket user ID is empty")
	}

	overwrites, err := ticketPermissionOverwrites(s, data.GuildID, data.UserID)
	if err != nil {
		return nil, err
	}

	channel, err := s.GuildChannelCreateComplex(data.GuildID, discordgo.GuildChannelCreateData{
		Name:                 ticketChannelName(s, data),
		Type:                 discordgo.ChannelTypeGuildText,
		PermissionOverwrites: overwrites,
		Topic:                fmt.Sprintf("Ticket %s | User ID: %s", data.Type, data.UserID),
	})
	if err != nil {
		return nil, fmt.Errorf("create ticket channel: %w", err)
	}

	if _, err = s.ChannelMessageSendComplex(
		channel.ID,
		&discordgo.MessageSend{
			Components: ticketsMessages.OpenTicketMessageComponents(data),
			Flags:      ticketsData.MessageFlagsIsComponentsV2,
		},
	); err != nil {
		log.Printf("failed to send start message to ticket channel %s: %v", channel.ID, err)
	}

	return channel, nil
}

func ticketPermissionOverwrites(s *discordgo.Session, guildID string, userID string) ([]*discordgo.PermissionOverwrite, error) {
	guild, err := s.Guild(guildID)
	if err != nil {
		return nil, fmt.Errorf("get guild %s: %w", guildID, err)
	}

	allowTicketMember := int64(discordgo.PermissionViewChannel |
		discordgo.PermissionSendMessages |
		discordgo.PermissionReadMessageHistory |
		discordgo.PermissionAttachFiles |
		discordgo.PermissionEmbedLinks)

	allowAdmin := int64(discordgo.PermissionViewChannel |
		discordgo.PermissionSendMessages |
		discordgo.PermissionReadMessageHistory |
		discordgo.PermissionManageChannels |
		discordgo.PermissionManageMessages)

	overwrites := []*discordgo.PermissionOverwrite{
		{
			ID:   guildID,
			Type: discordgo.PermissionOverwriteTypeRole,
			Deny: int64(discordgo.PermissionViewChannel),
		},
		{
			ID:    userID,
			Type:  discordgo.PermissionOverwriteTypeMember,
			Allow: allowTicketMember,
		},
	}

	if guild.OwnerID != "" && guild.OwnerID != userID {
		overwrites = append(overwrites, &discordgo.PermissionOverwrite{
			ID:    guild.OwnerID,
			Type:  discordgo.PermissionOverwriteTypeMember,
			Allow: allowAdmin,
		})
	}

	for _, role := range guild.Roles {
		if role.Permissions&discordgo.PermissionAdministrator == 0 {
			continue
		}

		overwrites = append(overwrites, &discordgo.PermissionOverwrite{
			ID:    role.ID,
			Type:  discordgo.PermissionOverwriteTypeRole,
			Allow: allowAdmin,
		})
	}

	return overwrites, nil
}

func ticketChannelName(s *discordgo.Session, data ticketsData.TicketChannelData) string {

	username := sanitizeChannelName(ticketCreatorName(s, data.GuildID, data.UserID))
	if username == "" {
		username = "user"
	}

	ticketType := sanitizeChannelName(string(data.Type))
	if ticketType == "" {
		ticketType = "Невідомо"
	}

	return fmt.Sprintf("-%s-%s", username, ticketType)
}

func ticketCreatorName(s *discordgo.Session, guildID string, userID string) string {
	member, err := s.GuildMember(guildID, userID)
	if err == nil && member != nil {
		if member.Nick != "" {
			return member.Nick
		}

		if member.User != nil {
			return member.User.Username
		}
	}

	user, err := s.User(userID)
	if err == nil && user != nil {
		return user.Username
	}

	return userID
}

func sanitizeChannelName(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	value = strings.ReplaceAll(value, " ", "-")

	re := regexp.MustCompile(`[^a-z0-9а-яіїєґ-]+`)
	value = re.ReplaceAllString(value, "")

	re = regexp.MustCompile(`-+`)
	return strings.Trim(re.ReplaceAllString(value, "-"), "-")
}
