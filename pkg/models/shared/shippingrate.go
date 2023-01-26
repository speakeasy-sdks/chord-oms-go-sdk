package shared

type ShippingRate struct {
	Cost               *string `json:"cost,omitempty"`
	DisplayCost        *string `json:"display_cost,omitempty"`
	ID                 *int64  `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	Selected           *bool   `json:"selected,omitempty"`
	ShippingMethodCode *string `json:"shipping_method_code,omitempty"`
	ShippingMethodID   *int64  `json:"shipping_method_id,omitempty"`
}
