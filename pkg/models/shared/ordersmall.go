package shared

import (
	"time"
)

type OrderSmall struct {
	AdditionalTaxTotal                      string                 `json:"additional_tax_total"`
	AdjustmentTotal                         string                 `json:"adjustment_total"`
	CancelerID                              *int64                 `json:"canceler_id,omitempty"`
	Channel                                 *string                `json:"channel,omitempty"`
	CheckoutSteps                           []string               `json:"checkout_steps"`
	CompletedAt                             time.Time              `json:"completed_at"`
	CoveredByStoreCredit                    *bool                  `json:"covered_by_store_credit,omitempty"`
	CreatedAt                               time.Time              `json:"created_at"`
	CrossSellDetails                        *CrossSellDetails      `json:"cross_sell_details,omitempty"`
	Currency                                string                 `json:"currency"`
	DisplayAdditionalTaxTotal               string                 `json:"display_additional_tax_total"`
	DisplayIncludedTaxTotal                 string                 `json:"display_included_tax_total"`
	DisplayItemTotal                        string                 `json:"display_item_total"`
	DisplayOrderTotalAfterStoreCredit       *string                `json:"display_order_total_after_store_credit,omitempty"`
	DisplayShipTotal                        string                 `json:"display_ship_total"`
	DisplayStoreCreditRemainingAfterCapture *string                `json:"display_store_credit_remaining_after_capture,omitempty"`
	DisplayTaxTotal                         *string                `json:"display_tax_total,omitempty"`
	DisplayTotal                            string                 `json:"display_total"`
	DisplayTotalApplicableStoreCredit       *string                `json:"display_total_applicable_store_credit,omitempty"`
	DisplayTotalAvailableStoreCredit        *string                `json:"display_total_available_store_credit,omitempty"`
	Email                                   string                 `json:"email"`
	ID                                      int64                  `json:"id"`
	IncludedTaxTotal                        string                 `json:"included_tax_total"`
	ItemTotal                               string                 `json:"item_total"`
	LineItems                               []LineItem             `json:"line_items,omitempty"`
	Locale                                  string                 `json:"locale"`
	Metadata                                map[string]interface{} `json:"metadata,omitempty"`
	Number                                  string                 `json:"number"`
	OrderTotalAfterStoreCredit              *string                `json:"order_total_after_store_credit,omitempty"`
	PaymentState                            *string                `json:"payment_state,omitempty"`
	PaymentTotal                            string                 `json:"payment_total"`
	ShipTotal                               string                 `json:"ship_total"`
	ShipmentState                           *string                `json:"shipment_state,omitempty"`
	SpecialInstructions                     *string                `json:"special_instructions,omitempty"`
	State                                   *string                `json:"state,omitempty"`
	TaxTotal                                string                 `json:"tax_total"`
	Token                                   string                 `json:"token"`
	Total                                   string                 `json:"total"`
	TotalApplicableStoreCredit              *interface{}           `json:"total_applicable_store_credit,omitempty"`
	TotalQuantity                           *int64                 `json:"total_quantity,omitempty"`
	UpdatedAt                               time.Time              `json:"updated_at"`
	UserID                                  int64                  `json:"user_id"`
}
