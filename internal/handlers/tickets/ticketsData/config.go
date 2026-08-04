package ticketsData

type TicketSettingsConfig struct {
	DefaultCategoryID  string         `yaml:"default_category_id"`
	DefaultChannelName string         `yaml:"default_channel_name"`
	ArchiveCategoryID  string         `yaml:"archive_category_id"`
	LogChannelID       string         `yaml:"log_channel_id"`
	DefaultStaffRoles  []string       `yaml:"default_staff_roles"`
	DefaultButtons     []TicketButton `yaml:"default_buttons"`
}

type TicketsConfigFile struct {
	TicketSettings TicketSettingsConfig    `yaml:"ticket_settings"`
	Tickets        map[string]TicketConfig `yaml:"tickets"`
}

type TicketConfig struct {
	Title       string         `yaml:"title"`
	CategoryID  string         `yaml:"category_id"`
	ChannelName string         `yaml:"channel_name"`
	StaffRoles  []string       `yaml:"staff_roles"`
	Modal       TicketModal    `yaml:"modal"`
	Message     TicketMessage  `yaml:"message"`
	Buttons     []TicketButton `yaml:"buttons"`
}

type TicketModal struct {
	Fields []TicketModalField `yaml:"fields"`
}

type TicketModalField struct {
	ID          string `yaml:"id"`
	Label       string `yaml:"label"`
	Style       string `yaml:"style"`
	Placeholder string `yaml:"placeholder"`
	Required    bool   `yaml:"required"`
	MinLength   int    `yaml:"min_length"`
	MaxLength   int    `yaml:"max_length"`
}

type TicketMessage struct {
	Ping        string `yaml:"ping"`
	Header      string `yaml:"header"`
	Description string `yaml:"description"`
	Color       string `yaml:"color"`
}

type TicketButton struct {
	ID    string `yaml:"id"`
	Label string `yaml:"label"`
	Style string `yaml:"style"`
	Emoji string `yaml:"emoji"`
}

type ResolvedTicketConfig struct {
	TicketConfig
	TypeKey string
}

func (cfg *TicketsConfigFile) Resolve(typeKey string) (ResolvedTicketConfig, bool) {
	ticket, ok := cfg.Tickets[typeKey]
	if !ok {
		return ResolvedTicketConfig{}, false
	}

	if ticket.CategoryID == "" {
		ticket.CategoryID = cfg.TicketSettings.DefaultCategoryID
	}
	if ticket.ChannelName == "" {
		ticket.ChannelName = cfg.TicketSettings.DefaultChannelName
	}
	if len(ticket.StaffRoles) == 0 {
		ticket.StaffRoles = cfg.TicketSettings.DefaultStaffRoles
	}
	if len(ticket.Buttons) == 0 {
		ticket.Buttons = cfg.TicketSettings.DefaultButtons
	}

	return ResolvedTicketConfig{TicketConfig: ticket, TypeKey: typeKey}, true
}
