package shared

import (
	"time"
)

type CrossSellDetails struct {
	CrossSellDurationInMinutes *int64     `json:"cross_sell_duration_in_minutes,omitempty"`
	CrossSellEndsAt            *time.Time `json:"cross_sell_ends_at,omitempty"`
	CrossSellable              *bool      `json:"cross_sellable,omitempty"`
}
