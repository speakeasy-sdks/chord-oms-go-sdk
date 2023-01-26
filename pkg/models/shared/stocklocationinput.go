package shared

type StockLocationInput struct {
	Active               *bool   `json:"active,omitempty"`
	Address1             *string `json:"address1,omitempty"`
	Address2             *string `json:"address2,omitempty"`
	BackorderableDefault *bool   `json:"backorderable_default,omitempty"`
	City                 *string `json:"city,omitempty"`
	CountryID            *int64  `json:"country_id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Phone                *string `json:"phone,omitempty"`
	PropagateAllVariants *bool   `json:"propagate_all_variants,omitempty"`
	StateID              *int64  `json:"state_id,omitempty"`
	StateName            *string `json:"state_name,omitempty"`
	Zipcode              *string `json:"zipcode,omitempty"`
}
