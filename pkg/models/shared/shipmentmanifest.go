package shared

type ShipmentManifestStates struct {
	OnHand *int64 `json:"on_hand,omitempty"`
}

type ShipmentManifest struct {
	Quantity  *int64                  `json:"quantity,omitempty"`
	Sku       *string                 `json:"sku,omitempty"`
	States    *ShipmentManifestStates `json:"states,omitempty"`
	VariantID *int64                  `json:"variant_id,omitempty"`
}
