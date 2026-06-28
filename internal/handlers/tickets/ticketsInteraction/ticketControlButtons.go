package ticketsInteraction

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"fmt"
	"log"

	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func TicketControlButtons(cfg config.Env) func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}

		data := i.MessageComponentData()
		switch data.CustomID {
		// case ticketsOpenLogic.TicketAcceptButtonID:
		//  handleTicketAccept(s, i, cfg)
		case ticketsData.TicketCloseButtonID:
			handleTicketClose(s, i, cfg)
		}
	}
}

/*
func handleTicketAccept(s *discordgo.Session, i *discordgo.InteractionCreate, cfg config.Env) {
	if !canUseTicketControl(s, i, cfg.TicketAcceptRoleIDs) {
		respondTicketControl(s, i, "❌ У вас немає прав прийняти тикет.", true)
		return
	}

	respondTicketControl(s, i, fmt.Sprintf("✅ Тикет прийняв %s.", ticketActorMention(i)), false)
}
*/

func handleTicketClose(s *discordgo.Session, i *discordgo.InteractionCreate, cfg config.Env) {
	if !canUseTicketControl(s, i, cfg.TicketCloseRoleIDs) {
		respondTicketControl(s, i, "❌ У вас немає прав закрити цю заявку.", true)
		return
	}

	respondTicketControl(s, i, fmt.Sprintf("🔒 Заявку закрив %s.", ticketActorMention(i)), false)

	if _, err := s.ChannelDelete(i.ChannelID); err != nil {
		log.Printf("failed to close ticket channel %s: %v", i.ChannelID, err)
	}
}

func canUseTicketControl(s *discordgo.Session, i *discordgo.InteractionCreate, roleIDs []string) bool {
	if i.Member == nil || i.Member.User == nil {
		return false
	}

	permissions, err := s.UserChannelPermissions(i.Member.User.ID, i.ChannelID)
	if err == nil && permissions&discordgo.PermissionAdministrator != 0 {
		return true
	}

	if len(roleIDs) == 0 {
		return false
	}

	allowedRoles := make(map[string]struct{}, len(roleIDs))
	for _, roleID := range roleIDs {
		if roleID == "" {
			continue
		}

		allowedRoles[roleID] = struct{}{}
	}

	for _, memberRoleID := range i.Member.Roles {
		if _, ok := allowedRoles[memberRoleID]; ok {
			return true
		}
	}

	return false
}

func ticketActorMention(i *discordgo.InteractionCreate) string {
	if i.Member != nil && i.Member.User != nil {
		return fmt.Sprintf("<@%s>", i.Member.User.ID)
	}

	if i.User != nil {
		return fmt.Sprintf("<@%s>", i.User.ID)
	}

	return "невідомий користувач"
}

func respondTicketControl(s *discordgo.Session, i *discordgo.InteractionCreate, content string, ephemeral bool) {
	flags := discordgo.MessageFlags(0)
	if ephemeral {
		flags = discordgo.MessageFlagsEphemeral
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   flags,
		},
	}); err != nil {
		log.Printf("failed to respond ticket control interaction: %v", err)
	}
}
