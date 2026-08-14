package cats

import (
	"STVbotGo/internal/backend"
	"context"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func showCat(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	catsNames := map[string]struct{}{
		"котик":   {},
		"котики":  {},
		"кот":     {},
		"коты":    {},
		"кошка":   {},
		"кошки":   {},
		"кошечка": {},
		"кошечки": {},
	}
	content := strings.TrimSpace(strings.ToLower(m.Content))
	if _, ok := catsNames[content]; ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := backend.NewClient(
		"https://api.thecatapi.com",
		"",
	)

	var cat string

	cat, err := client.RandomCat(ctx)
	if err != nil || cat == "" {
		log.Println("Не вдалося отримати зображення кота:", err)
		cat = "Не вдалося отримати зображення кота"
	}

	_, err = s.ChannelMessageSendReply(m.ChannelID, cat, m.Reference())
	if err != nil {
		log.Println(err)
		return
	}
}
