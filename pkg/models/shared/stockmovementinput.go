package shared

type StockMovementInput struct {
	Quantity    *int64 `json:"quantity,omitempty"`
	StockItemID *int64 `json:"stock_item_id,omitempty"`
}
