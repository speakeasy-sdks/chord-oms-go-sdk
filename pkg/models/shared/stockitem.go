package shared

type StockItem struct {
	Backorderable   *bool    `json:"backorderable,omitempty"`
	CountOnHand     *int64   `json:"count_on_hand,omitempty"`
	ID              *int64   `json:"id,omitempty"`
	StockLocationID *int64   `json:"stock_location_id,omitempty"`
	Variant         *Variant `json:"variant,omitempty"`
	VariantID       *int64   `json:"variant_id,omitempty"`
}
