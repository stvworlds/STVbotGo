package ticketsOpenLogic

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"STVbotGo/internal/logger"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func CreateTicketChannel(s *discordgo.Session, data ticketsData.TicketChannelData) (*discordgo.Channel, error) {
	if data.GuildID == "" {
		return nil, fmt.Errorf("guild ID is empty")
	}

	if data.UserID == "" {
		return nil, fmt.Errorf("ticket user ID is empty")
	}

	cfg, ok := ticketsData.Config.Resolve(data.Type)
	if !ok {
		return nil, fmt.Errorf("unknown ticket type: %s", data.Type)
	}

	overwrites, err := ticketPermissionOverwrites(s, data.GuildID, data.UserID)
	if err != nil {
		return nil, err
	}

	username := sanitizeChannelName(ticketCreatorName(s, data.GuildID, data.UserID))
	if username == "" {
		username = "user"
	}

	templateVars := buildTemplateVars(data, username)
	channelName := sanitizeChannelName(ticketsData.RenderTemplate(cfg.ChannelName, templateVars))

	channel, err := s.GuildChannelCreateComplex(data.GuildID, discordgo.GuildChannelCreateData{
		Name:                 channelName,
		Type:                 discordgo.ChannelTypeGuildText,
		ParentID:             cfg.CategoryID,
		PermissionOverwrites: overwrites,
		Topic:                fmt.Sprintf("Ticket %s | User ID: %s", data.Type, data.UserID),
	})
	if err != nil {
		return nil, fmt.Errorf("create ticket channel: %w", err)
	}

	avatarURL := ticketCreatorAvatar(s, data.GuildID, data.UserID)
	components := ticketsData.BuildTicketMessage(cfg.Message, cfg.Buttons, templateVars, avatarURL)
	if _, err = s.ChannelMessageSendComplex(
		channel.ID,
		&discordgo.MessageSend{
			Components: components,
			Flags:      ticketsData.MessageFlagsIsComponentsV2,
		},
	); err != nil {
		logger.WithError(err, fmt.Sprintf("failed to send start message to ticket channel %s", channel.ID))
	}

	return channel, nil
}

func buildTemplateVars(data ticketsData.TicketChannelData, username string) map[string]string {
	vars := map[string]string{
		"user_id":      data.UserID,
		"username":     username,
		"user_mention": fmt.Sprintf("<@%s>", data.UserID),
		"guild_id":     data.GuildID,
		"ticket_type":  data.Type,
		"timestamp":    fmt.Sprintf("%d", time.Now().Unix()),
	}
	for k, v := range data.Fields {
		vars[k] = v
	}
	return vars
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

func ticketCreatorAvatar(s *discordgo.Session, guildID string, userID string) string {
	member, err := s.GuildMember(guildID, userID)
	if err == nil && member != nil && member.User != nil {
		return member.User.AvatarURL("256")
	}

	user, err := s.User(userID)
	if err == nil && user != nil {
		return user.AvatarURL("256")
	}

	return ""
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
