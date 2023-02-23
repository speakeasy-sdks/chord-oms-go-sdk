package shared

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/types"
	"time"
)

type User struct {
	AvailableGiftCards  []GiftCard             `json:"available_gift_cards,omitempty"`
	BillAddress         *Address               `json:"bill_address,omitempty"`
	CreatedAt           *time.Time             `json:"created_at,omitempty"`
	DateOfBirth         *types.Date            `json:"date_of_birth,omitempty"`
	Email               *string                `json:"email,omitempty"`
	ID                  *int64                 `json:"id,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
	Name                *string                `json:"name,omitempty"`
	Notes               *string                `json:"notes,omitempty"`
	Phone               *string                `json:"phone,omitempty"`
	Roles               []string               `json:"roles,omitempty"`
	ShipAddress         *Address               `json:"ship_address,omitempty"`
	StoreCredit         *string                `json:"store_credit,omitempty"`
	TrackingPreferences []TrackingPreference   `json:"tracking_preferences,omitempty"`
	UpdatedAt           *time.Time             `json:"updated_at,omitempty"`
}
