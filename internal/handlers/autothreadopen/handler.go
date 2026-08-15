package autothreadopen

import (
	"STVbotGo/internal/logger"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (h *cfg) autoThreadOpen(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(h.newsChannels) == 0 {
		logger.Info("News channels configuration is missing. Auto thread opening is disabled.")
		return
	}

	if _, ok := h.newsChannels[m.ChannelID]; !ok {
		return
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to get channel %s", m.ChannelID))
		return
	}

	if channel.IsThread() {
		return
	}

	if _, err = s.ChannelMessageCrosspost(m.ChannelID, m.ID); err != nil {
		logger.WithError(err, fmt.Sprintf("failed to crosspost message %s in channel %s", m.ID, m.ChannelID))
		return
	}

	thread, err := s.MessageThreadStart(m.ChannelID, m.ID, ".", 60)
	if err != nil {
		logger.WithError(err, fmt.Sprintf("failed to create thread for message %s in channel %s", m.ID, m.ChannelID))
		return
	}

	if _, err = s.ChannelMessageSend(thread.ID, "Обговорення 🔽"); err != nil {
		logger.WithError(err, fmt.Sprintf("failed to send message to thread %s", thread.ID))
	}
}
