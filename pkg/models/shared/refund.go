package shared

type Refund struct {
	Amount          *int64  `json:"amount,omitempty"`
	CreatedAt       *string `json:"created_at,omitempty"`
	ID              *int64  `json:"id,omitempty"`
	PaymentID       *int64  `json:"payment_id,omitempty"`
	RefundReason    *string `json:"refund_reason,omitempty"`
	ReimbursementID *int64  `json:"reimbursement_id,omitempty"`
	TransactionID   *string `json:"transaction_id,omitempty"`
	UpdatedAt       *string `json:"updated_at,omitempty"`
}
