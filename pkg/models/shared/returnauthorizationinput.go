package shared

type ReturnAuthorizationInputReturnItemsAttributes struct {
	ExchangeVariantID *int64 `json:"exchange_variant_id,omitempty"`
	InventoryUnitID   *int64 `json:"inventory_unit_id,omitempty"`
	ReturnReasonID    *int64 `json:"return_reason_id,omitempty"`
}

type ReturnAuthorizationInput struct {
	Memo                  *string                                         `json:"memo,omitempty"`
	ReturnItemsAttributes []ReturnAuthorizationInputReturnItemsAttributes `json:"return_items_attributes,omitempty"`
	ReturnReasonID        *int64                                          `json:"return_reason_id,omitempty"`
	StockLocationID       *int64                                          `json:"stock_location_id,omitempty"`
}
