package tickets

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func TicketSelectMenu(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	data := i.MessageComponentData()

	if data.CustomID != "stv_tickets_select" {
		return
	}
	if len(data.Values) == 0 {
		return
	}

	selected := data.Values[0]

	var response []discordgo.MessageComponent

	switch selected {
	case "pidtrimka":
		response = pidtrimkaComponents()
	case "skargi":
		response = skargiComponents()
	default:
		response = defaultComponents()
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Components: response,
			Flags:      messageFlagsIsComponentsV2,
		},
	})
	if err != nil {
		log.Printf("Error responding to select menu interaction: %v", err)
		return
	}
}

func skargiComponents() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{

				discordgo.TextDisplay{
					Content: "нажміте кнопку чтобі аткрить тікет в скаргі",
				},

				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Аткрить",
							Style:    discordgo.SuccessButton,
							Disabled: false,
							CustomID: "stv_tickets_open_skargi",
						},
					},
				},
			},
		},
	}
}

func defaultComponents() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{

				discordgo.TextDisplay{
					Content: "разраб гавнакодер",
				},
			},
		},
	}
}

func pidtrimkaComponents() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{

				discordgo.TextDisplay{
					Content: "нажміте кнопку чтобі аткрить тікет в паджерку",
				},

				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Аткрить",
							Style:    discordgo.SuccessButton,
							Disabled: false,
							CustomID: "stv_tickets_open_pidtrimka",
						},
					},
				},
			},
		},
	}
}
