package shared

type Payment struct {
	Amount          *string        `json:"amount,omitempty"`
	AvsResponse     *string        `json:"avs_response,omitempty"`
	CreatedAt       *string        `json:"created_at,omitempty"`
	DisplayAmount   *string        `json:"display_amount,omitempty"`
	ID              *int64         `json:"id,omitempty"`
	PaymentMethod   *PaymentMethod `json:"payment_method,omitempty"`
	PaymentMethodID *int64         `json:"payment_method_id,omitempty"`
	Source          *PaymentSource `json:"source,omitempty"`
	SourceID        *int64         `json:"source_id,omitempty"`
	SourceType      *int64         `json:"source_type,omitempty"`
	State           *string        `json:"state,omitempty"`
	UpdatedAt       *string        `json:"updated_at,omitempty"`
}
