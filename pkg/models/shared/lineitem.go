package shared

type LineItem struct {
	Adjustments         []Adjustment           `json:"adjustments"`
	CrossSold           *int64                 `json:"cross_sold,omitempty"`
	DisplayAmount       *string                `json:"display_amount,omitempty"`
	DynamicPromotions   []DynamicPromotion     `json:"dynamic_promotions,omitempty"`
	GiftCards           []GiftCard             `json:"gift_cards,omitempty"`
	ID                  *int64                 `json:"id,omitempty"`
	IsBundle            *bool                  `json:"is_bundle,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
	Parts               []Part                 `json:"parts,omitempty"`
	Price               *string                `json:"price,omitempty"`
	Quantity            *int64                 `json:"quantity,omitempty"`
	SingleDisplayAmount *string                `json:"single_display_amount,omitempty"`
	SubscriptionType    *string                `json:"subscription_type,omitempty"`
	Total               *string                `json:"total,omitempty"`
	Variant             *Variant               `json:"variant,omitempty"`
	VariantID           *int64                 `json:"variant_id,omitempty"`
}
