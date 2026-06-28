package ticketsData

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	// TicketCloseButtonID TicketAcceptButtonID = "stv_tickets_accept"
	TicketCloseButtonID                               = "stv_tickets_close"
	MessageFlagsIsComponentsV2 discordgo.MessageFlags = 1 << 15
)

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

type TicketType string

const (
	TicketSupport     TicketType = "support"
	TicketAdminReport TicketType = "admin-report"
	TicketUserReport  TicketType = "user-report"
	TicketMessage     TicketType = "message"
)

type ModalField struct {
	CustomID    string
	Label       string
	Style       discordgo.TextInputStyle
	Placeholder string
	Required    bool
	MaxLength   int
}

type TicketFormConfig struct {
	Type   TicketType
	Title  string
	Fields []ModalField
}
