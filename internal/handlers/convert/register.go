package convert

import "github.com/bwmarrin/discordgo"

const (
	convertCommandName = "convert"
	filterComplex      = "[0:v]fps=15,scale=360:-1:flags=lanczos,split[s0][s1];[s0]palettegen=max_colors=256:stats_mode=diff[p];[s1][p]paletteuse=dither=sierra2_4a:diff_mode=rectangle"
)

func Register(s *discordgo.Session) {
	s.AddHandler(commandConvert)
}

func RegisterCommand(s *discordgo.Session, applicationID string) error {
	_, err := s.ApplicationCommandCreate(applicationID, "", &discordgo.ApplicationCommand{
		Name:        convertCommandName,
		Description: "Convert file to GIF",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionAttachment,
				Name:        "file",
				Description: "Input file for conversion",
				Required:    true,
			},
		},
	})
	return err
}
