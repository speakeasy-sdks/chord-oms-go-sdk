package shared

type SubscriptionLineItemIntervalUnitsEnum string

const (
	SubscriptionLineItemIntervalUnitsEnumDay   SubscriptionLineItemIntervalUnitsEnum = "day"
	SubscriptionLineItemIntervalUnitsEnumWeek  SubscriptionLineItemIntervalUnitsEnum = "week"
	SubscriptionLineItemIntervalUnitsEnumMonth SubscriptionLineItemIntervalUnitsEnum = "month"
)

type SubscriptionLineItem struct {
	Destroy          *bool                                  `json:"_destroy,omitempty"`
	Description      *string                                `json:"description,omitempty"`
	DisplayPrice     *string                                `json:"display_price,omitempty"`
	EndDate          *string                                `json:"end_date,omitempty"`
	ID               *float64                               `json:"id,omitempty"`
	Images           []Image                                `json:"images,omitempty"`
	Interval         *string                                `json:"interval,omitempty"`
	IntervalLength   *float64                               `json:"interval_length,omitempty"`
	IntervalUnits    *SubscriptionLineItemIntervalUnitsEnum `json:"interval_units,omitempty"`
	Name             *string                                `json:"name,omitempty"`
	OptionsText      *string                                `json:"options_text,omitempty"`
	Price            *string                                `json:"price,omitempty"`
	ProductID        *float64                               `json:"product_id,omitempty"`
	ProductSlug      *string                                `json:"product_slug,omitempty"`
	Quantity         *float64                               `json:"quantity,omitempty"`
	Sku              *string                                `json:"sku,omitempty"`
	SubscriptionType *string                                `json:"subscription_type,omitempty"`
	VariantID        *float64                               `json:"variant_id,omitempty"`
}
