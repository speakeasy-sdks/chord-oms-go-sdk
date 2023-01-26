package shared

type ProductProperty struct {
	ID           *int64  `json:"id,omitempty"`
	ProductID    *int64  `json:"product_id,omitempty"`
	PropertyID   *int64  `json:"property_id,omitempty"`
	PropertyName *string `json:"property_name,omitempty"`
	Value        *string `json:"value,omitempty"`
}
