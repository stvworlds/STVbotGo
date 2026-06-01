package autothreadopen

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *cfg) autoThreadOpen(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(h.newsChannels) == 0 {
		log.Println("News channels configuration is missing. Auto thread opening is disabled.")
		return
	}

	if _, ok := h.newsChannels[m.ChannelID]; !ok {
		return
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Printf("Failed to get channel %s: %v", m.ChannelID, err)
		return
	}

	if channel.IsThread() {
		return
	}

	if _, err = s.ChannelMessageCrosspost(m.ChannelID, m.ID); err != nil {
		log.Printf("Failed to crosspost message %s in channel %s: %v", m.ID, m.ChannelID, err)
		return
	}

	thread, err := s.MessageThreadStart(m.ChannelID, m.ID, ".", 60)
	if err != nil {
		log.Printf("Failed to create thread for message %s in channel %s: %v", m.ID, m.ChannelID, err)
		return
	}

	if _, err = s.ChannelMessageSend(thread.ID, "Обговорення 🔽"); err != nil {
		log.Printf("Failed to send message to thread %s: %v", thread.ID, err)
	}
}
