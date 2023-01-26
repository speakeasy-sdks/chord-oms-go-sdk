package shared

type Adjustment struct {
	AdjustableID    *int64       `json:"adjustable_id,omitempty"`
	AdjustableType  *string      `json:"adjustable_type,omitempty"`
	Amount          *string      `json:"amount,omitempty"`
	CreatedAt       *string      `json:"created_at,omitempty"`
	DisplayAmount   *string      `json:"display_amount,omitempty"`
	Eligible        *bool        `json:"eligible,omitempty"`
	Finalized       *bool        `json:"finalized,omitempty"`
	ID              *int64       `json:"id,omitempty"`
	Label           *string      `json:"label,omitempty"`
	PromotionCode   *interface{} `json:"promotion_code,omitempty"`
	PromotionCodeID *int64       `json:"promotion_code_id,omitempty"`
	SourceID        *int64       `json:"source_id,omitempty"`
	SourceType      *string      `json:"source_type,omitempty"`
	UpdatedAt       *string      `json:"updated_at,omitempty"`
}
