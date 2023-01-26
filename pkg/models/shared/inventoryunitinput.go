package shared

type InventoryUnitInput struct {
	Shipment  *string `json:"shipment,omitempty"`
	VariantID *int64  `json:"variant_id,omitempty"`
}
