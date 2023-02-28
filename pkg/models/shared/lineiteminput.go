package shared

// LineItemInputOptions
// This field can be used to pass custom line item attributes. When used, it will force a new price calculation, unless `price` is one of the options.
type LineItemInputOptions struct {
	Parts []Part `json:"parts,omitempty"`
}

type LineItemInput struct {
	DynamicPromotionsAttributes     []DynamicPromotion     `json:"dynamic_promotions_attributes,omitempty"`
	Options                         *LineItemInputOptions  `json:"options,omitempty"`
	Parts                           []Part                 `json:"parts,omitempty"`
	Quantity                        *int64                 `json:"quantity,omitempty"`
	SubscriptionLineItemsAttributes []SubscriptionLineItem `json:"subscription_line_items_attributes,omitempty"`
}
