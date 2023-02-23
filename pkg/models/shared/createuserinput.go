package shared

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/types"
)

type CreateUserInputUser struct {
	BillAddressAttributes *AddressInput          `json:"bill_address_attributes,omitempty"`
	DateOfBirth           *types.Date            `json:"date_of_birth,omitempty"`
	Email                 string                 `json:"email"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	Notes                 *string                `json:"notes,omitempty"`
	Password              string                 `json:"password"`
	PasswordConfirmation  string                 `json:"password_confirmation"`
	Phone                 *string                `json:"phone,omitempty"`
	ShipAddressAttributes *AddressInput          `json:"ship_address_attributes,omitempty"`
}

type CreateUserInput struct {
	User CreateUserInputUser `json:"user"`
}
