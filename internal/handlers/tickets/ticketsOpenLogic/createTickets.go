package ticketsOpenLogic

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type TicketCreateData struct {
	Type    ticketsData.TicketType
	UserID  string
	GuildID string
	Fields  map[string]string
}

func CreateTickets(s *discordgo.Session, i *discordgo.InteractionCreate, data TicketCreateData) {
	logTicketCreateData(data)

	channel, err := CreateTicketChannel(s, ticketsData.TicketChannelData{
		Type:    data.Type,
		UserID:  data.UserID,
		GuildID: data.GuildID,
		Fields:  data.Fields,
	})
	if err != nil {
		log.Printf("failed to create %s ticket channel: %v", data.Type, err)
		respondTicketCreateResult(s, i, "❌ Не вдалося створити канал заявки. Напишіть лист до <support@stvworlds.win>.")
		return
	}

	respondTicketCreateResult(s, i, fmt.Sprintf("✅ Заявку створено, <#%s>.", channel.ID))
}

func respondTicketCreateResult(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("failed to send ticket create response: %v", err)
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
