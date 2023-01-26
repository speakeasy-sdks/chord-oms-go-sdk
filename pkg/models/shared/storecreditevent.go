package shared

type StoreCreditEvent struct {
	DisplayAction          *string `json:"display_action,omitempty"`
	DisplayAmount          *string `json:"display_amount,omitempty"`
	DisplayEventDate       *string `json:"display_event_date,omitempty"`
	DisplayRemainingAmount *string `json:"display_remaining_amount,omitempty"`
	DisplayUserTotalAmount *string `json:"display_user_total_amount,omitempty"`
	OrderNumber            *string `json:"order_number,omitempty"`
}
