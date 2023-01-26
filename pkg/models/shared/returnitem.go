package shared

type ReturnItem struct {
	AcceptanceStatus           *string `json:"acceptance_status,omitempty"`
	AcceptanceStatusErrors     *string `json:"acceptance_status_errors,omitempty"`
	AdditionalTaxTotal         *int64  `json:"additional_tax_total,omitempty"`
	Amount                     *int64  `json:"amount,omitempty"`
	CreatedAt                  *string `json:"created_at,omitempty"`
	CustomerReturn             *string `json:"customer_return,omitempty"`
	ExchangeInventoryUnitID    *int64  `json:"exchange_inventory_unit_id,omitempty"`
	ExchangeVariantID          *int64  `json:"exchange_variant_id,omitempty"`
	ID                         *int64  `json:"id,omitempty"`
	IncludedTaxTotal           *int64  `json:"included_tax_total,omitempty"`
	InventoryUnitID            *int64  `json:"inventory_unit_id,omitempty"`
	OverrideReimbursementType  *string `json:"override_reimbursement_type,omitempty"`
	PreferredReimbursementType *string `json:"preferred_reimbursement_type,omitempty"`
	ReceptionStatus            *string `json:"reception_status,omitempty"`
	Reimbursement              *string `json:"reimbursement,omitempty"`
	Resellable                 *bool   `json:"resellable,omitempty"`
	UpdatedAt                  *string `json:"updated_at,omitempty"`
}
