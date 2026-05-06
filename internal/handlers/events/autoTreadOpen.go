package events

import (
	"github.com/bwmarrin/discordgo"
)

func AutoThreadOpen(s *discordgo.Session, m *discordgo.MessageCreate) {

	Id := make([]int, 0)
	Id = append(Id, 1501679320937005136)
	for i := 0; i < len(Id); i++ {
		if m.ChannelID != Id[i] {
			return
		}
	}
	if m.ChanelType != 5 {
		return
	}
	s.MessageThreadStart(m.ChannelID, m.ID, "Auto Thread")
}
