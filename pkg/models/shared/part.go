package shared

type Part struct {
	Quantity *int64  `json:"quantity,omitempty"`
	Sku      *string `json:"sku,omitempty"`
}
