package shared

type ProductInput struct {
	AvailableOn                 *string                `json:"available_on,omitempty"`
	CostCurrency                *string                `json:"cost_currency,omitempty"`
	CostPrice                   *string                `json:"cost_price,omitempty"`
	DeletedAt                   *string                `json:"deleted_at,omitempty"`
	Depth                       *string                `json:"depth,omitempty"`
	Description                 *string                `json:"description,omitempty"`
	Height                      *string                `json:"height,omitempty"`
	MetaDescription             *string                `json:"meta_description,omitempty"`
	MetaKeywords                *string                `json:"meta_keywords,omitempty"`
	Name                        *string                `json:"name,omitempty"`
	OptionTypeIds               []int64                `json:"option_type_ids,omitempty"`
	OptionValuesHash            map[string]interface{} `json:"option_values_hash,omitempty"`
	Permalink                   *string                `json:"permalink,omitempty"`
	Price                       *string                `json:"price,omitempty"`
	ProductPropertiesAttributes []ProductPropertyInput `json:"product_properties_attributes,omitempty"`
	ShippingCategoryID          *int64                 `json:"shipping_category_id,omitempty"`
	Sku                         *string                `json:"sku,omitempty"`
	TaxCategoryID               *int64                 `json:"tax_category_id,omitempty"`
	TaxonIds                    []int64                `json:"taxon_ids,omitempty"`
	Weight                      *string                `json:"weight,omitempty"`
	Width                       *string                `json:"width,omitempty"`
}
