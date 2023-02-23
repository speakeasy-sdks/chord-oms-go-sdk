package shared

type AddressBookInput struct {
	Address1         string  `json:"address1"`
	Address2         *string `json:"address2,omitempty"`
	AlternativePhone *string `json:"alternative_phone,omitempty"`
	City             string  `json:"city"`
	Company          *string `json:"company,omitempty"`
	CountryID        *int64  `json:"country_id,omitempty"`
	CountryIso       *string `json:"country_iso,omitempty"`
	Default          *bool   `json:"default,omitempty"`
	Firstname        *string `json:"firstname,omitempty"`
	Lastname         *string `json:"lastname,omitempty"`
	Name             *string `json:"name,omitempty"`
	Phone            *string `json:"phone,omitempty"`
	StateID          *int64  `json:"state_id,omitempty"`
	StateName        *string `json:"state_name,omitempty"`
	Zipcode          *string `json:"zipcode,omitempty"`
}
