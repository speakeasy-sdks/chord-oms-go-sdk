package shared

type ShippingMethod struct {
	Code               *string            `json:"code,omitempty"`
	ID                 *int64             `json:"id,omitempty"`
	Name               *string            `json:"name,omitempty"`
	ShippingCategories []ShippingCategory `json:"shipping_categories,omitempty"`
	Zones              []Zone             `json:"zones,omitempty"`
}
