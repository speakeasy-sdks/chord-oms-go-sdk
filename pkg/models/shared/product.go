package shared

import (
	"time"
)

type Product struct {
	AvailableOn        *time.Time        `json:"available_on,omitempty"`
	Classifications    []Classification  `json:"classifications,omitempty"`
	Description        *string           `json:"description,omitempty"`
	DisplayPrice       *string           `json:"display_price,omitempty"`
	HasVariants        *bool             `json:"has_variants,omitempty"`
	ID                 *int64            `json:"id,omitempty"`
	Master             *Variant          `json:"master,omitempty"`
	MetaDescription    *string           `json:"meta_description,omitempty"`
	MetaKeywords       *string           `json:"meta_keywords,omitempty"`
	MetaTitle          *string           `json:"meta_title,omitempty"`
	Name               *string           `json:"name,omitempty"`
	OptionTypes        []OptionType      `json:"option_types,omitempty"`
	Price              *string           `json:"price,omitempty"`
	ProductProperties  []ProductProperty `json:"product_properties,omitempty"`
	ShippingCategoryID *int64            `json:"shipping_category_id,omitempty"`
	Slug               *string           `json:"slug,omitempty"`
	TaxonIds           []int64           `json:"taxon_ids,omitempty"`
	TotalOnHand        *int64            `json:"total_on_hand,omitempty"`
	Variants           []Variant         `json:"variants,omitempty"`
}
