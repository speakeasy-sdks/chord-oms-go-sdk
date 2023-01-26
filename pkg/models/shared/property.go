package shared

type Property struct {
	ID           *int64  `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Presentation *string `json:"presentation,omitempty"`
}
