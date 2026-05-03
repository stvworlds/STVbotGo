package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)
import "STVbotGo/internal/handlers/events"

func eventsLoader(s *discordgo.Session) {

	s.AddHandler(events.OnReady)
	s.AddHandler(events.OnMemberJoin)
	s.AddHandler(events.OnIpMessage)

	log.Println("Events are now loaded.")
}
