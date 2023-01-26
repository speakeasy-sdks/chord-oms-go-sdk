package shared

type CreditCardUpdateInput struct {
	AddressAttributes *AddressInput `json:"address_attributes,omitempty"`
	Expirty           *string       `json:"expirty,omitempty"`
	FirstName         *string       `json:"first_name,omitempty"`
	LastName          *string       `json:"last_name,omitempty"`
	Month             *string       `json:"month,omitempty"`
	Name              *string       `json:"name,omitempty"`
	Year              *string       `json:"year,omitempty"`
}
