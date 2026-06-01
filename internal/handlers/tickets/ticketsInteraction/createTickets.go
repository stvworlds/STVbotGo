package ticketsInteraction

import (
	"fmt"
	"log"
	"strings"

	"STVbotGo/internal/handlers/tickets/ticketsForms"

	"github.com/bwmarrin/discordgo"
)

type TicketCreateData struct {
	Type    ticketsForms.TicketType
	UserID  string
	GuildID string
	Fields  map[string]string
}

func CreateTickets(s *discordgo.Session, i *discordgo.InteractionCreate, data TicketCreateData) {
	logTicketCreateData(data)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: buildTicketCreatedMessage(data),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("failed to create %s ticket: %v", data.Type, err)
	}
}

func logTicketCreateData(data TicketCreateData) {
	fields := make([]string, 0, len(data.Fields))

	for key, value := range data.Fields {
		fields = append(fields, fmt.Sprintf("%s=%q", key, value))
	}

	log.Printf(
		"ticket created: type=%q userID=%q guildID=%q fields={%s}",
		data.Type,
		data.UserID,
		data.GuildID,
		strings.Join(fields, " "),
	)
}

func buildTicketCreatedMessage(data TicketCreateData) string {
	return "✅ Ваше звернення отримано. Обробник поки працює як заглушка."
}
