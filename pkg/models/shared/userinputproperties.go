package shared

import (
	"time"
)

type UserInputProperties struct {
	BillAddressAttributes *AddressInput          `json:"bill_address_attributes,omitempty"`
	DateOfBirth           *time.Time             `json:"date_of_birth,omitempty"`
	Email                 *string                `json:"email,omitempty"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	Notes                 *string                `json:"notes,omitempty"`
	Password              *string                `json:"password,omitempty"`
	PasswordConfirmation  *string                `json:"password_confirmation,omitempty"`
	Phone                 *string                `json:"phone,omitempty"`
	ShipAddressAttributes *AddressInput          `json:"ship_address_attributes,omitempty"`
}
