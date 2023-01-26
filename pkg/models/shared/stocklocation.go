package shared

type StockLocation struct {
	Active    *bool    `json:"active,omitempty"`
	Address1  *string  `json:"address1,omitempty"`
	Address2  *string  `json:"address2,omitempty"`
	City      *string  `json:"city,omitempty"`
	Country   *Country `json:"country,omitempty"`
	CountryID *int64   `json:"country_id,omitempty"`
	ID        *int64   `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Phone     *string  `json:"phone,omitempty"`
	State     *State   `json:"state,omitempty"`
	StateID   *int64   `json:"state_id,omitempty"`
	StateName *string  `json:"state_name,omitempty"`
	Zipcode   *string  `json:"zipcode,omitempty"`
}
