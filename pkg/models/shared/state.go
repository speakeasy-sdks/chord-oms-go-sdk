package shared

type State struct {
	Abbr      *string `json:"abbr,omitempty"`
	CountryID *int64  `json:"country_id,omitempty"`
	ID        *int64  `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
}
