package ticketsForms

import (
	"STVbotGo/internal/handlers/tickets/ticketsData"
	"log"

	"github.com/bwmarrin/discordgo"
)

const (
	TicketModalPrefix = "stv_tickets_"
	TicketModalSuffix = "_modal"

	TicketButtonPrefix = "stv_tickets_open_"
)

var ticketForms = map[ticketsData.TicketType]ticketsData.TicketFormConfig{
	ticketsData.TicketSupport: {
		Type:  ticketsData.TicketSupport,
		Title: "Звернення до підтримки",
		Fields: []ticketsData.ModalField{
			{
				CustomID:    "topic",
				Label:       "Тема звернення",
				Style:       discordgo.TextInputShort,
				Placeholder: "Наприклад: проблема з входом на сервер",
				Required:    true,
				MaxLength:   100,
			},
			{
				CustomID:    "description",
				Label:       "Опишіть проблему",
				Style:       discordgo.TextInputParagraph,
				Placeholder: "Опишіть вашу проблему якомога детальніше",
				Required:    true,
				MaxLength:   1000,
			},
		},
	},
	ticketsData.TicketAdminReport: {
		Type:  ticketsData.TicketAdminReport,
		Title: "Скарга на адміністрацію",
		Fields: []ticketsData.ModalField{
			{
				CustomID:    "topic",
				Label:       "Тема скарги",
				Style:       discordgo.TextInputShort,
				Placeholder: "Наприклад: перевищення обов'язків адміністратором",
				Required:    true,
				MaxLength:   100,
			},
			{
				CustomID:    "description",
				Label:       "Опишіть скаргу",
				Style:       discordgo.TextInputParagraph,
				Placeholder: "Опишіть вашу скаргу якомога детальніше",
				Required:    true,
				MaxLength:   1000,
			},
		},
	},
	ticketsData.TicketUserReport: {
		Type:  ticketsData.TicketUserReport,
		Title: "Скарга на гравця",
		Fields: []ticketsData.ModalField{
			{
				CustomID:    "topic",
				Label:       "Тема скарги",
				Style:       discordgo.TextInputShort,
				Placeholder: "Наприклад: образи від іншого гравця в вашу сторону",
				Required:    true,
				MaxLength:   100,
			},
			{
				CustomID:    "description",
				Label:       "Опишіть скаргу",
				Style:       discordgo.TextInputParagraph,
				Placeholder: "Опишіть вашу скаргу якомога детальніше",
				Required:    true,
				MaxLength:   1000,
			},
		},
	},
	ticketsData.TicketMessage: {
		Type:  ticketsData.TicketMessage,
		Title: "Повідомлення дла адміністрації",
		Fields: []ticketsData.ModalField{
			{
				CustomID:    "topic",
				Label:       "Тема повідомлення",
				Style:       discordgo.TextInputShort,
				Placeholder: "Наприклад: ідея для нового івенту.",
				Required:    true,
				MaxLength:   100,
			},
			{
				CustomID:    "description",
				Label:       "Опишіть скаргу",
				Style:       discordgo.TextInputParagraph,
				Placeholder: "Опишіть ваше повідомлення якомога детальніше",
				Required:    true,
				MaxLength:   1000,
			},
		},
	},
}

func ModalID(ticketType ticketsData.TicketType) string {
	return TicketModalPrefix + string(ticketType) + TicketModalSuffix
}

func ButtonID(ticketType ticketsData.TicketType) string {
	return TicketButtonPrefix + string(ticketType)
}

func IsKnownForm(ticketType ticketsData.TicketType) bool {
	_, ok := ticketForms[ticketType]
	return ok
}

func OpenTicketForm(s *discordgo.Session, i *discordgo.InteractionCreate, ticketType ticketsData.TicketType) {
	form, ok := ticketForms[ticketType]
	if !ok {
		log.Printf("unknown ticket form type: %s", ticketType)
		return
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID:   ModalID(ticketType),
			Title:      form.Title,
			Components: buildModalComponents(form.Fields),
		},
	})
	if err != nil {
		log.Printf("failed to open %s ticket modal: %v", ticketType, err)
	}
}

func buildModalComponents(fields []ticketsData.ModalField) []discordgo.MessageComponent {
	components := make([]discordgo.MessageComponent, 0, len(fields))

	for _, field := range fields {
		components = append(components, discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.TextInput{
					CustomID:    field.CustomID,
					Label:       field.Label,
					Style:       field.Style,
					Placeholder: field.Placeholder,
					Required:    field.Required,
					MaxLength:   field.MaxLength,
				},
			},
		})
	}

	return components
}
