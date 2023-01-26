package shared

// LineItemInputOptions
// This field can be used to pass custom line item attributes. When used, it will force a new price calculation, unless `price` is one of the options.
type LineItemInputOptions struct {
	Parts []Part `json:"parts,omitempty"`
}

type LineItemInput struct {
	DynamicPromotionsAttributes     []DynamicPromotion     `json:"dynamic_promotions_attributes,omitempty"`
	ID                              *int64                 `json:"id,omitempty"`
	Options                         *LineItemInputOptions  `json:"options,omitempty"`
	Parts                           []Part                 `json:"parts,omitempty"`
	Quantity                        *int64                 `json:"quantity,omitempty"`
	Sku                             *string                `json:"sku,omitempty"`
	SubscriptionLineItemsAttributes []SubscriptionLineItem `json:"subscription_line_items_attributes,omitempty"`
	VariantID                       *int64                 `json:"variant_id,omitempty"`
}
