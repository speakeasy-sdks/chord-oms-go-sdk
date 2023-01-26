package shared

type StockItemInput struct {
	Backorderable *bool  `json:"backorderable,omitempty"`
	StockLocation *int64 `json:"stock_location,omitempty"`
	Variant       *int64 `json:"variant,omitempty"`
	VariantID     *int64 `json:"variant_id,omitempty"`
}
