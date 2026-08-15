package newscreate

import (
	"STVbotGo/internal/backend"
	"STVbotGo/internal/config"
	"STVbotGo/internal/logger"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func extractNewsID(responseBody []byte) (string, error) {
	var response backend.CreateNewsResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "", fmt.Errorf("parse create news response: %w", err)
	}

	if response.Status != "success" {
		return "", fmt.Errorf("api returned unsuccessful status: %s", response.Status)
	}

	if response.Data.ID <= 0 {
		return "", fmt.Errorf("api response does not contain valid news id")
	}

	return strconv.Itoa(response.Data.ID), nil
}

func checkAccessRole(s *discordgo.Session, i *discordgo.InteractionCreate, cfg config.Env) bool {
	roleID := cfg.NewsRoleID
	member, err := s.GuildMember(cfg.NewsGuildID, i.Member.User.ID)
	if err != nil {
		logger.WithError(err, "failed to fetch guild member")
		return false
	}

	for _, r := range member.Roles {
		if r == roleID {
			return true
		}
	}

	return false
}

func checkUserHasTimeout(s *discordgo.Session, i *discordgo.InteractionCreate, cfg config.Env) bool {
	member, err := s.GuildMember(cfg.NewsGuildID, i.Member.User.ID)
	if err != nil {
		logger.WithError(err, "failed to fetch guild member")
		return false
	}
	if member.CommunicationDisabledUntil != nil {
		return true
	}

	return false
}

func sanitizeNewsInput(value string) string {
	value = strings.TrimSpace(value)

	replacer := strings.NewReplacer(
		`"`, "",
		`'`, "",
		`;`, "",
	)

	value = replacer.Replace(value)

	return strings.TrimSpace(value)
}

func isAllowedThumbnailURL(value string) bool {
	value = strings.ToLower(strings.TrimSpace(value))

	return strings.HasPrefix(value, "https://") || strings.HasPrefix(value, "http://")
}

func getInteractionUserID(i *discordgo.InteractionCreate) string {
	if i.Member != nil && i.Member.User != nil {
		return i.Member.User.ID
	}

	if i.User != nil {
		return i.User.ID
	}

	return ""
}

func getModalValue(i *discordgo.InteractionCreate, customID string) string {
	for _, component := range i.ModalSubmitData().Components {
		row, ok := component.(*discordgo.ActionsRow)
		if !ok {
			continue
		}

		for _, rowComponent := range row.Components {
			input, ok := rowComponent.(*discordgo.TextInput)
			if !ok {
				continue
			}

			if input.CustomID == customID {
				return strings.TrimSpace(input.Value)
			}
		}
	}

	return ""
}

func respondEphemeral(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		logger.WithError(err, "failed to send ephemeral response")
	}
}
