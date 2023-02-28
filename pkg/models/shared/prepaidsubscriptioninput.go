package shared

type PrepaidSubscriptionInputInterval struct {
	Length float64          `json:"length"`
	Unit   IntervalUnitEnum `json:"unit"`
}

type PrepaidSubscriptionInputLineItems struct {
	Quantity float64 `json:"quantity"`
	Sku      string  `json:"sku"`
}

// PrepaidSubscriptionInput
// The input to add a prepaid gift subscription to an order.
//
// ### Notes
//
//   - If `gift: false`, then you do not need to provide `recipient` or
//     of `ship_address`. The subscription will be shipped to the address
//     entered during checkout.
//   - If `gift: true`, you must provide the aforementioned fields before
//     as a part of the request as they cannot be added or updated later.
type PrepaidSubscriptionInput struct {
	AutoRedeem       bool                                `json:"auto_redeem"`
	InstallmentCount float64                             `json:"installment_count"`
	Interval         PrepaidSubscriptionInputInterval    `json:"interval"`
	LineItems        []PrepaidSubscriptionInputLineItems `json:"line_items"`
}
