package shared

type OrderBig struct {
	AdditionalTaxTotal                      *string                `json:"additional_tax_total,omitempty"`
	AdjustmentTotal                         *string                `json:"adjustment_total,omitempty"`
	Adjustments                             []Adjustment           `json:"adjustments,omitempty"`
	BillAddress                             *Address               `json:"bill_address,omitempty"`
	CancelerID                              *int64                 `json:"canceler_id,omitempty"`
	Channel                                 *string                `json:"channel,omitempty"`
	CheckoutSteps                           []string               `json:"checkout_steps,omitempty"`
	CompletedAt                             *string                `json:"completed_at,omitempty"`
	CoveredByStoreCredit                    *bool                  `json:"covered_by_store_credit,omitempty"`
	CreatedAt                               *string                `json:"created_at,omitempty"`
	CreditCards                             []CreditCard           `json:"credit_cards,omitempty"`
	CrossSellDurationInMinutes              *int64                 `json:"cross_sell_duration_in_minutes,omitempty"`
	CrossSellEndsAt                         *string                `json:"cross_sell_ends_at,omitempty"`
	CrossSellable                           *bool                  `json:"cross_sellable,omitempty"`
	Currency                                *string                `json:"currency,omitempty"`
	DisplayAdditionalTaxTotal               *string                `json:"display_additional_tax_total,omitempty"`
	DisplayIncludedTaxTotal                 *string                `json:"display_included_tax_total,omitempty"`
	DisplayItemTotal                        *string                `json:"display_item_total,omitempty"`
	DisplayOrderTotalAfterStoreCredit       *string                `json:"display_order_total_after_store_credit,omitempty"`
	DisplayShipTotal                        *string                `json:"display_ship_total,omitempty"`
	DisplayStoreCreditRemainingAfterCapture *string                `json:"display_store_credit_remaining_after_capture,omitempty"`
	DisplayTaxTotal                         *string                `json:"display_tax_total,omitempty"`
	DisplayTotal                            *string                `json:"display_total,omitempty"`
	DisplayTotalApplicableStoreCredit       *string                `json:"display_total_applicable_store_credit,omitempty"`
	DisplayTotalAvailableStoreCredit        *string                `json:"display_total_available_store_credit,omitempty"`
	Email                                   *string                `json:"email,omitempty"`
	GiftNote                                *GiftNote              `json:"gift_note,omitempty"`
	ID                                      *int64                 `json:"id,omitempty"`
	IncludedTaxTotal                        *string                `json:"included_tax_total,omitempty"`
	ItemTotal                               *string                `json:"item_total,omitempty"`
	LineItems                               []LineItem             `json:"line_items,omitempty"`
	Locale                                  *string                `json:"locale,omitempty"`
	Metadata                                map[string]interface{} `json:"metadata,omitempty"`
	Number                                  *string                `json:"number,omitempty"`
	OrderTotalAfterStoreCredit              *string                `json:"order_total_after_store_credit,omitempty"`
	PaymentMethods                          []PaymentMethod        `json:"payment_methods,omitempty"`
	PaymentState                            *string                `json:"payment_state,omitempty"`
	PaymentTotal                            *string                `json:"payment_total,omitempty"`
	Payments                                []Payment              `json:"payments,omitempty"`
	Permissions                             *OrderPermissions      `json:"permissions,omitempty"`
	ShipAddress                             *Address               `json:"ship_address,omitempty"`
	ShipTotal                               *string                `json:"ship_total,omitempty"`
	ShipmentState                           *string                `json:"shipment_state,omitempty"`
	Shipments                               []Shipment             `json:"shipments,omitempty"`
	SpecialInstructions                     *string                `json:"special_instructions,omitempty"`
	State                                   *string                `json:"state,omitempty"`
	TagList                                 []string               `json:"tag_list,omitempty"`
	TaxTotal                                *string                `json:"tax_total,omitempty"`
	Token                                   *string                `json:"token,omitempty"`
	Total                                   *string                `json:"total,omitempty"`
	TotalApplicableStoreCredit              *string                `json:"total_applicable_store_credit,omitempty"`
	TotalQuantity                           *int64                 `json:"total_quantity,omitempty"`
	UpdatedAt                               *string                `json:"updated_at,omitempty"`
	UserID                                  *int64                 `json:"user_id,omitempty"`
}