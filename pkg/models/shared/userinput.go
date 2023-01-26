package shared

type UserInputUser struct {
	BillAddressAttributes *AddressInput `json:"bill_address_attributes,omitempty"`
	Email                 *string       `json:"email,omitempty"`
	Password              *string       `json:"password,omitempty"`
	PasswordConfirmation  *string       `json:"password_confirmation,omitempty"`
	ShipAddressAttributes *AddressInput `json:"ship_address_attributes,omitempty"`
}

type UserInput struct {
	User *UserInputUser `json:"user,omitempty"`
}
