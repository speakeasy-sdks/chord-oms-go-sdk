package shared

type InventoryUnit struct {
	ID         *int64  `json:"id,omitempty"`
	ShipmentID *int64  `json:"shipment_id,omitempty"`
	State      *string `json:"state,omitempty"`
	VariantID  *int64  `json:"variant_id,omitempty"`
}
