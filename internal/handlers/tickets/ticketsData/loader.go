package ticketsData

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Config *TicketsConfigFile

func LoadTicketsConfig(path string) (*TicketsConfigFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read tickets config %s: %w", path, err)
	}

	var cfg TicketsConfigFile
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse tickets config %s: %w", path, err)
	}

	if err = ValidateTicketsConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ValidateTicketsConfig(cfg *TicketsConfigFile) error {
	if len(cfg.Tickets) == 0 {
		return fmt.Errorf("tickets config: no ticket types defined")
	}

	validButtonStyles := map[string]bool{
		"primary": true, "secondary": true, "success": true, "danger": true,
	}

	for name, ticket := range cfg.Tickets {
		if ticket.Title == "" {
			return fmt.Errorf("ticket %q: title is required", name)
		}

		if len(ticket.Modal.Fields) == 0 {
			return fmt.Errorf("ticket %q: modal.fields is required", name)
		}

		if len(ticket.Modal.Fields) > 5 {
			return fmt.Errorf("ticket %q: modal.fields cannot exceed 5 (Discord limit)", name)
		}

		for i, field := range ticket.Modal.Fields {
			if field.ID == "" {
				return fmt.Errorf("ticket %q field[%d]: id is required", name, i)
			}
			if field.Label == "" {
				return fmt.Errorf("ticket %q field[%d]: label is required", name, i)
			}
			if field.Style != "short" && field.Style != "paragraph" {
				return fmt.Errorf("ticket %q field[%d]: style must be 'short' or 'paragraph', got %q", name, i, field.Style)
			}
		}

		for i, btn := range ticket.Buttons {
			if !validButtonStyles[btn.Style] {
				return fmt.Errorf("ticket %q button[%d]: style must be primary/secondary/success/danger, got %q", name, i, btn.Style)
			}
		}

		r, _ := cfg.Resolve(name)
		if r.CategoryID == "" {
			return fmt.Errorf("ticket %q: category_id is required (no default set)", name)
		}
		if r.ChannelName == "" {
			return fmt.Errorf("ticket %q: channel_name is required (no default set)", name)
		}
	}

	return nil
}
