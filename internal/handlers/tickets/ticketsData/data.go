package ticketsData

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	TicketCloseButtonID                               = "stv_tickets_close"
	MessageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
)

type TicketType = string
type TicketChannelData struct {
	Type    TicketType
	UserID  string
	GuildID string
	Fields  map[string]string
}

func Ptr[T any](value T) *T {
	return &value
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var Rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[Rng.Intn(len(chars))]
	}
	return string(result)
}
