package events

import (
	"STVbotGo/internal/config"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	register(func(s *discordgo.Session, cfg config.Env) {
		AntiScamConfig := antiSpamHandler(cfg.ScamChannelID, cfg.ActionsLogChannelID)
		s.AddHandler(AntiScamConfig.antiSpam)
	})
}

type AntiScamConfig struct {
	ScamChannelID       string `json:"1510035959033499709"`
	ActionsLogChannelID string
}

func antiSpamHandler(scamChannelID string, actionsLogChannelID string) *AntiScamConfig {
	return &AntiScamConfig{
		ScamChannelID:       scamChannelID,
		ActionsLogChannelID: actionsLogChannelID,
	}
}

const (
	deleteMessageSeconds = 60 * 60
)

func (h *AntiScamConfig) antiSpam(s *discordgo.Session, m *discordgo.MessageCreate) {

	banStatus := "Ok"

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Message.ChannelID != h.ScamChannelID {
		return
	}

	if len(m.Attachments) < 2 {
		return
	}

	msgSend := &discordgo.MessageSend{
		Reference: m.Message.Forward(),
	}
	_, err := s.ChannelMessageSendComplex(h.ActionsLogChannelID, msgSend)
	if err != nil {
		log.Printf("Failed to send message to actions log channel: %v", err)
		return
	}

	if err := banUser(s, m); err != nil {
		banStatus = "Failed"
	}

	sendAdminMessage(s, m, h.ActionsLogChannelID, banStatus)
}

func banUser(s *discordgo.Session, m *discordgo.MessageCreate) error {
	data := struct {
		DeleteMessageSeconds int `json:"delete_message_seconds"`
	}{
		DeleteMessageSeconds: deleteMessageSeconds,
	}

	_, err := s.RequestWithBucketID(
		"PUT",
		discordgo.EndpointGuildBan(m.GuildID, m.Author.ID),
		data,
		discordgo.EndpointGuildBan(m.GuildID, ""),
		discordgo.WithAuditLogReason("Anti-scam triggered"),
	)
	if err != nil {
		log.Printf("Failed to ban user for anti-scam: %v", err)
		return err
	}

	return nil
}

func sendAdminMessage(s *discordgo.Session, m *discordgo.MessageCreate, channel string, status string) {
	_, err := s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Components: adminMessageComponents(m, status),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Println("Failed to send admin message:", err)
		return
	}
}

func adminMessageComponents(m *discordgo.MessageCreate, status string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{
				discordgo.Section{
					Components: []discordgo.MessageComponent{
						discordgo.TextDisplay{
							Content: m.Author.Mention() + " був відмічений системою як спамер. \nCтатус покарання: " + status,
						},
					},
					Accessory: discordgo.Thumbnail{
						Media: discordgo.UnfurledMediaItem{
							URL: m.Author.AvatarURL("256"),
						},
						Description: ptr("Scam Avatar"),
					},
				},
			},
		},
		discordgo.Container{
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: fmt.Sprintf("📜 Айді користувача: \n```\n%s\n```", m.Author.ID),
				},
			},
		},
		discordgo.Container{
			AccentColor: ptr(6164128),
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: fmt.Sprintf("🎰 Час приєднання до серверу:\n<t:%d:R>", m.Member.JoinedAt.Unix()),
				},
			},
		},
	}
}
