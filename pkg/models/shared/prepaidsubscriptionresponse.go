package shared

type PrepaidSubscriptionResponseInterval struct {
	Length float64          `json:"length"`
	Unit   IntervalUnitEnum `json:"unit"`
}

type PrepaidSubscriptionResponseLineItems struct {
	Quantity float64 `json:"quantity"`
	Sku      string  `json:"sku"`
}

type PrepaidSubscriptionResponseOptionsLineItems struct {
	Price    string `json:"price"`
	Quantity int64  `json:"quantity"`
	Sku      string `json:"sku"`
}

type PrepaidSubscriptionResponseOptions struct {
	LineItems []PrepaidSubscriptionResponseOptionsLineItems `json:"line_items,omitempty"`
}

type PrepaidSubscriptionResponseStateEnum string

const (
	PrepaidSubscriptionResponseStateEnumPending   PrepaidSubscriptionResponseStateEnum = "pending"
	PrepaidSubscriptionResponseStateEnumActive    PrepaidSubscriptionResponseStateEnum = "active"
	PrepaidSubscriptionResponseStateEnumCompleted PrepaidSubscriptionResponseStateEnum = "completed"
)

// PrepaidSubscriptionResponse
// The input to add a prepaid gift subscription to an order.
//
// ### Notes
//
//   - If `gift: false`, then you do not need to provide `recipient` or
//     of `ship_address`. The subscription will be shipped to the address
//     entered during checkout.
//   - If `gift: true`, you must provide the aforementioned fields before
//     as a part of the request as they cannot be added or updated later.
type PrepaidSubscriptionResponse struct {
	AddressID              int64                                  `json:"address_id"`
	Amount                 string                                 `json:"amount"`
	AutoRedeem             bool                                   `json:"auto_redeem"`
	ID                     int64                                  `json:"id"`
	InstallmentCount       float64                                `json:"installment_count"`
	Interval               PrepaidSubscriptionResponseInterval    `json:"interval"`
	LineItemID             int64                                  `json:"line_item_id"`
	LineItems              []PrepaidSubscriptionResponseLineItems `json:"line_items"`
	Options                PrepaidSubscriptionResponseOptions     `json:"options"`
	Quantity               int64                                  `json:"quantity"`
	SingleInstallmentPrice string                                 `json:"single_installment_price"`
	State                  PrepaidSubscriptionResponseStateEnum   `json:"state"`
	SubscriptionID         int64                                  `json:"subscription_id"`
	VirtualGiftCardID      int64                                  `json:"virtual_gift_card_id"`
}
