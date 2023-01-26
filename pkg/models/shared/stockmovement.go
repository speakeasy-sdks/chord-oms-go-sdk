package shared

type StockMovement struct {
	ID          *int64     `json:"id,omitempty"`
	Quantity    *int64     `json:"quantity,omitempty"`
	StockItem   *StockItem `json:"stock_item,omitempty"`
	StockItemID *int64     `json:"stock_item_id,omitempty"`
}
