package shared

type OptionTypeInput struct {
	Name                   *string            `json:"name,omitempty"`
	OptionValuesAttributes []OptionValueInput `json:"option_values_attributes,omitempty"`
	Presentation           *string            `json:"presentation,omitempty"`
}
