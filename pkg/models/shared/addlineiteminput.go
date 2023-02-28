package shared

type AddLineItemInputOptions struct {
	Parts               []Part                    `json:"parts,omitempty"`
	PrepaidSubscription *PrepaidSubscriptionInput `json:"prepaid_subscription,omitempty"`
}

type AddLineItemInput struct {
	DynamicPromotionsAttributes     []DynamicPromotion       `json:"dynamic_promotions_attributes,omitempty"`
	Options                         *AddLineItemInputOptions `json:"options,omitempty"`
	Parts                           []Part                   `json:"parts,omitempty"`
	Quantity                        *int64                   `json:"quantity,omitempty"`
	SubscriptionLineItemsAttributes []SubscriptionLineItem   `json:"subscription_line_items_attributes,omitempty"`
}
