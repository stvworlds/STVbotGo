package ticketsInteraction

import (
	"STVbotGo/internal/tickets/ticketsMessages"
	"log"

	"github.com/bwmarrin/discordgo"
)

const (
	messageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
)

func TicketSelectMenu(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type != discordgo.InteractionMessageComponent {
		return
	}
	if i.MessageComponentData().CustomID != "stv_tickets_select" {
		return
	}
	if len(i.MessageComponentData().Values) == 0 {
		return
	}

	selected := i.MessageComponentData().Values[0]
	var response []discordgo.MessageComponent

	switch selected {
	case "support", "admin-report", "user-report", "message":
		response = ticketsMessages.OpenBottonComponents(selected)
	default:
		response = ticketsMessages.DefaultErrorComponents()
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Components: response,
			Flags:      messageFlagsIsComponentsV2 | discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("Error responding to select menu interaction: %v", err)
		return
	}
}
