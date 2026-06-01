package ticketsMessages

import "github.com/bwmarrin/discordgo"

func DefaultErrorComponents() []discordgo.MessageComponent {
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
