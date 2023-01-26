package shared

type AddressBook struct {
	Address1         *string  `json:"address1,omitempty"`
	Address2         *string  `json:"address2,omitempty"`
	AlternativePhone *string  `json:"alternative_phone,omitempty"`
	City             *string  `json:"city,omitempty"`
	Company          *string  `json:"company,omitempty"`
	Country          *Country `json:"country,omitempty"`
	CountryID        *int64   `json:"country_id,omitempty"`
	CountryIso       *string  `json:"country_iso,omitempty"`
	Default          *bool    `json:"default,omitempty"`
	ID               *int64   `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Phone            *string  `json:"phone,omitempty"`
	State            *State   `json:"state,omitempty"`
	StateID          *int64   `json:"state_id,omitempty"`
	StateName        *string  `json:"state_name,omitempty"`
	StateText        *string  `json:"state_text,omitempty"`
	Zipcode          *string  `json:"zipcode,omitempty"`
}
