package shared

type AddressInputCountry struct {
	Iso     *string `json:"iso,omitempty"`
	Iso3    *string `json:"iso3,omitempty"`
	IsoName *string `json:"iso_name,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type AddressInputState struct {
	Abbr *string `json:"abbr,omitempty"`
	Name *string `json:"name,omitempty"`
}

type AddressInput struct {
	Address1         *string              `json:"address1,omitempty"`
	Address2         *string              `json:"address2,omitempty"`
	AlternativePhone *string              `json:"alternative_phone,omitempty"`
	City             *string              `json:"city,omitempty"`
	Company          *string              `json:"company,omitempty"`
	Country          *AddressInputCountry `json:"country,omitempty"`
	CountryID        *int64               `json:"country_id,omitempty"`
	CountryIso       *string              `json:"country_iso,omitempty"`
	Firstname        *string              `json:"firstname,omitempty"`
	Lastname         *string              `json:"lastname,omitempty"`
	Name             *string              `json:"name,omitempty"`
	Phone            *string              `json:"phone,omitempty"`
	State            *AddressInputState   `json:"state,omitempty"`
	StateID          *int64               `json:"state_id,omitempty"`
	StateName        *string              `json:"state_name,omitempty"`
	Zipcode          *string              `json:"zipcode,omitempty"`
}
