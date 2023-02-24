package shared

import (
	"time"
)

type SubscriptionInputPaymentSourceAttributes struct {
	CcType                  *string `json:"cc_type,omitempty"`
	GatewayPaymentProfileID *string `json:"gateway_payment_profile_id,omitempty"`
	LastDigits              *string `json:"last_digits,omitempty"`
	Month                   *string `json:"month,omitempty"`
	Name                    *string `json:"name,omitempty"`
	Nonce                   *string `json:"nonce,omitempty"`
	PaymentMethodType       *string `json:"payment_method_type,omitempty"`
	Year                    *string `json:"year,omitempty"`
}

type SubscriptionInput struct {
	ActionableDate          *time.Time                                `json:"actionable_date,omitempty"`
	BillAddressAttributes   *AddressInput                             `json:"bill_address_attributes,omitempty"`
	EndDate                 *string                                   `json:"end_date,omitempty"`
	IntervalLength          *float64                                  `json:"interval_length,omitempty"`
	IntervalUnits           *IntervalUnitEnum                         `json:"interval_units,omitempty"`
	LineItemsAttributes     []SubscriptionLineItem                    `json:"line_items_attributes,omitempty"`
	PaymentSourceAttributes *SubscriptionInputPaymentSourceAttributes `json:"payment_source_attributes,omitempty"`
	ShipAddressAttributes   *AddressInput                             `json:"ship_address_attributes,omitempty"`
}
