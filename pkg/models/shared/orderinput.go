package shared

type OrderInput struct {
	BillAddressAttributes *AddressInput          `json:"bill_address_attributes,omitempty"`
	Channel               *string                `json:"channel,omitempty"`
	CouponCode            *string                `json:"coupon_code,omitempty"`
	Email                 *string                `json:"email,omitempty"`
	GiftNoteAttributes    *GiftNote              `json:"gift_note_attributes,omitempty"`
	LineItemsAttributes   []LineItemInput        `json:"line_items_attributes,omitempty"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	PaymentsAttributes    []PaymentInput         `json:"payments_attributes,omitempty"`
	ReferralIdentifier    *string                `json:"referral_identifier,omitempty"`
	ShipAddressAttributes *AddressInput          `json:"ship_address_attributes,omitempty"`
	ShipmentsAttributes   []ShipmentInput        `json:"shipments_attributes,omitempty"`
	SpecialInstructions   *string                `json:"special_instructions,omitempty"`
	UseBilling            *bool                  `json:"use_billing,omitempty"`
}
