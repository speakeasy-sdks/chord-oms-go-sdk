package shared

type PaymentMethod struct {
	ID          *int64  `json:"id,omitempty"`
	MethodType  *string `json:"method_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	PartialName *string `json:"partial_name,omitempty"`
}
