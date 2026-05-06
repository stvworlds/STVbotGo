package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type AutoThreadOpenHandler struct {
	NewsChannels map[string]struct{}
}

func NewAutoThreadOpenHandler(newsChannels []string) *AutoThreadOpenHandler {
	channels := make(map[string]struct{}, len(newsChannels))

	for _, channelID := range newsChannels {
		channels[channelID] = struct{}{}
	}

	return &AutoThreadOpenHandler{
		NewsChannels: channels,
	}
}

func (h *AutoThreadOpenHandler) AutoThreadOpen(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(h.NewsChannels) == 0 {
		log.Println("News channels configuration is missing. Auto thread opening is disabled.")
		return
	}

	if _, ok := h.NewsChannels[m.ChannelID]; !ok {
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
