package shared

type ProductPropertyInput struct {
	Position     *int64  `json:"position,omitempty"`
	PropertyName *string `json:"property_name,omitempty"`
	Value        *string `json:"value,omitempty"`
}
