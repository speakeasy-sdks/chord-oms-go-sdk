package shared

type ShipmentInput struct {
	ID                     *int64  `json:"id,omitempty"`
	SelectedShippingRateID *int64  `json:"selected_shipping_rate_id,omitempty"`
	SpecialInstructions    *string `json:"special_instructions,omitempty"`
	StockLocationID        *int64  `json:"stock_location_id,omitempty"`
	Tracking               *string `json:"tracking,omitempty"`
}
