package events

import (
	"fmt"

	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func init() {
	register(func(s *discordgo.Session, _ config.Env) {
		s.AddHandler(OnIpMessage)
	})
}

func OnIpMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content != "ping" {
		return
	}

	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		fmt.Println("error creating channel:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}
	// Then we send the message through the channel we created.
	_, err = s.ChannelMessageSend(channel.ID, "Pong!")
	if err != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		fmt.Println("error sending DM message:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
	}

}
