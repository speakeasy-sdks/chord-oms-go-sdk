package shared

type OptionType struct {
	ID           *int64        `json:"id,omitempty"`
	Name         *string       `json:"name,omitempty"`
	OptionValues []OptionValue `json:"option_values,omitempty"`
	Position     *int64        `json:"position,omitempty"`
	Presentation *string       `json:"presentation,omitempty"`
}
