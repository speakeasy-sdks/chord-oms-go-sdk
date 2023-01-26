package shared

type OptionValue struct {
	ID                     *int64  `json:"id,omitempty"`
	Name                   *string `json:"name,omitempty"`
	OptionTypeID           *int64  `json:"option_type_id,omitempty"`
	OptionTypeName         *string `json:"option_type_name,omitempty"`
	OptionTypePresentation *string `json:"option_type_presentation,omitempty"`
	Presentation           *string `json:"presentation,omitempty"`
}
