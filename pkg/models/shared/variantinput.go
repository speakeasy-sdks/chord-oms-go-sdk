package shared

type VariantInputOptions struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type VariantInput struct {
	CostCurrency           *string                `json:"cost_currency,omitempty"`
	CostPrice              *string                `json:"cost_price,omitempty"`
	Depth                  *string                `json:"depth,omitempty"`
	Height                 *string                `json:"height,omitempty"`
	LockVersion            *string                `json:"lock_version,omitempty"`
	Metadata               map[string]interface{} `json:"metadata,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	OptionValueIds         []int64                `json:"option_value_ids,omitempty"`
	OptionValuesAttributes []OptionValueInput     `json:"option_values_attributes,omitempty"`
	Options                *VariantInputOptions   `json:"options,omitempty"`
	Position               *int64                 `json:"position,omitempty"`
	Presentation           *string                `json:"presentation,omitempty"`
	Price                  *string                `json:"price,omitempty"`
	Product                *int64                 `json:"product,omitempty"`
	ProductID              *int64                 `json:"product_id,omitempty"`
	Sku                    *string                `json:"sku,omitempty"`
	TrackInventory         *bool                  `json:"track_inventory,omitempty"`
	Weight                 *string                `json:"weight,omitempty"`
	Width                  *string                `json:"width,omitempty"`
}
