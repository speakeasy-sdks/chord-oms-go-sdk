package shared

type Variant struct {
	CostPrice       *string                `json:"cost_price,omitempty"`
	Depth           *string                `json:"depth,omitempty"`
	Description     *string                `json:"description,omitempty"`
	DisplayPrice    *string                `json:"display_price,omitempty"`
	Height          *string                `json:"height,omitempty"`
	ID              *int64                 `json:"id,omitempty"`
	Images          []Image                `json:"images,omitempty"`
	InStock         *bool                  `json:"in_stock,omitempty"`
	IsBackorderable *bool                  `json:"is_backorderable,omitempty"`
	IsDestroyed     *bool                  `json:"is_destroyed,omitempty"`
	IsMaster        *bool                  `json:"is_master,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	Name            *string                `json:"name,omitempty"`
	OptionValues    []OptionValue          `json:"option_values,omitempty"`
	OptionsText     *string                `json:"options_text,omitempty"`
	Price           *string                `json:"price,omitempty"`
	Sku             *string                `json:"sku,omitempty"`
	Slug            *string                `json:"slug,omitempty"`
	TotalOnHand     *int64                 `json:"total_on_hand,omitempty"`
	TrackInventory  *bool                  `json:"track_inventory,omitempty"`
	Weight          *string                `json:"weight,omitempty"`
	Width           *string                `json:"width,omitempty"`
}
