package ticketsInteraction

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"STVbotGo/internal/handlers/tickets/ticketsMessages"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func TicketSelectMenu(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type != discordgo.InteractionMessageComponent {
		return
	}
	if !strings.Contains(i.MessageComponentData().CustomID, "stv_tickets_select") {
		return
	}
	if len(i.MessageComponentData().Values) == 0 {
		return
	}

	selected := i.MessageComponentData().Values[0]
	var response []discordgo.MessageComponent

	if ticketsData.Config != nil {
		if _, ok := ticketsData.Config.Tickets[selected]; ok {
			response = ticketsMessages.OpenBottonComponents(selected)
		} else {
			response = ticketsMessages.DefaultErrorComponents()
		}
	} else {
		response = ticketsMessages.DefaultErrorComponents()
	}

	resendMessage(s, i)

	_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Components: response,
		Flags:      ticketsData.MessageFlagsIsComponentsV2 | discordgo.MessageFlagsEphemeral,
	})
	if err != nil {
		log.Printf("Error responding to select menu interaction: %v", err)
		return
	}
}

func resendMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Components: ticketsMessages.OpenMessageComponents(),
			Flags:      ticketsData.MessageFlagsIsComponentsV2,
		},
	})
	if err != nil {
		log.Printf("Error updating select menu interaction message: %v", err)
	}
}
