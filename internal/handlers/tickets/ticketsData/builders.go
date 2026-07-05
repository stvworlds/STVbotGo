package ticketsData

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func BuildModalComponents(fields []TicketModalField) []discordgo.MessageComponent {
	components := make([]discordgo.MessageComponent, 0, len(fields))
	for _, field := range fields {
		components = append(components, discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.TextInput{
					CustomID:    field.ID,
					Label:       field.Label,
					Style:       parseTextInputStyle(field.Style),
					Placeholder: field.Placeholder,
					Required:    field.Required,
					MaxLength:   field.MaxLength,
				},
			},
		})
	}
	return components
}

func BuildButtonsRow(buttons []TicketButton) discordgo.ActionsRow {
	components := make([]discordgo.MessageComponent, 0, len(buttons))
	for _, btn := range buttons {
		b := discordgo.Button{
			CustomID: resolveButtonCustomID(btn.ID),
			Label:    btn.Label,
			Style:    parseButtonStyle(btn.Style),
		}
		if btn.Emoji != "" {
			b.Emoji = &discordgo.ComponentEmoji{Name: btn.Emoji}
		}
		components = append(components, b)
	}
	return discordgo.ActionsRow{Components: components}
}

func BuildTicketMessage(msg TicketMessage, buttons []TicketButton, values map[string]string, avatarURL string) []discordgo.MessageComponent {
	color := 0
	if msg.Color != "" {
		parsed, err := strconv.ParseInt(msg.Color, 16, 32)
		if err == nil {
			color = int(parsed)
		}
	}

	header := RenderTemplate(msg.Header, values)
	description := RenderTemplate(msg.Description, values)

	var inner []discordgo.MessageComponent

	if header != "" {
		var headerComponent discordgo.MessageComponent
		if avatarURL != "" {
			headerComponent = discordgo.Section{
				Components: []discordgo.MessageComponent{
					discordgo.TextDisplay{Content: header},
				},
				Accessory: discordgo.Thumbnail{
					Media: discordgo.UnfurledMediaItem{URL: avatarURL},
				},
			}
		} else {
			headerComponent = discordgo.TextDisplay{Content: header}
		}
		inner = append(inner, headerComponent)

		if description != "" {
			spacing := discordgo.SeparatorSpacingSizeLarge
			inner = append(inner,
				discordgo.Separator{Divider: Ptr(true), Spacing: &spacing},
				discordgo.TextDisplay{Content: description},
			)
		}
	} else {
		var textComponent discordgo.MessageComponent
		if avatarURL != "" {
			textComponent = discordgo.Section{
				Components: []discordgo.MessageComponent{
					discordgo.TextDisplay{Content: description},
				},
				Accessory: discordgo.Thumbnail{
					Media: discordgo.UnfurledMediaItem{URL: avatarURL},
				},
			}
		} else {
			textComponent = discordgo.TextDisplay{Content: description}
		}
		inner = append(inner, textComponent)
	}

	var result []discordgo.MessageComponent

	if msg.Ping != "" {
		if ping := RenderTemplate(msg.Ping, values); ping != "" {
			result = append(result, discordgo.TextDisplay{Content: ping})
		}
	}

	result = append(result, discordgo.Container{
		AccentColor: Ptr(color),
		Components:  inner,
	})

	if len(buttons) > 0 {
		result = append(result, BuildButtonsRow(buttons))
	}

	return result
}

func parseTextInputStyle(style string) discordgo.TextInputStyle {
	if strings.ToLower(style) == "paragraph" {
		return discordgo.TextInputParagraph
	}
	return discordgo.TextInputShort
}

func parseButtonStyle(style string) discordgo.ButtonStyle {
	switch strings.ToLower(style) {
	case "primary":
		return discordgo.PrimaryButton
	case "secondary":
		return discordgo.SecondaryButton
	case "success":
		return discordgo.SuccessButton
	case "danger":
		return discordgo.DangerButton
	default:
		return discordgo.SecondaryButton
	}
}

func resolveButtonCustomID(id string) string {
	if id == "close" {
		return TicketCloseButtonID
	}
	return "stv_tickets_btn_" + id
}
