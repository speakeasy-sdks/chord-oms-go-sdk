package shared

type GiftCard struct {
	Amount         *string `json:"amount,omitempty"`
	AutoRedeem     *bool   `json:"auto_redeem,omitempty"`
	Filled         *bool   `json:"filled,omitempty"`
	GiftMessage    *string `json:"gift_message,omitempty"`
	ID             *int64  `json:"id,omitempty"`
	PurchaserName  *string `json:"purchaser_name,omitempty"`
	RecipientEmail *string `json:"recipient_email,omitempty"`
	RecipientName  *string `json:"recipient_name,omitempty"`
	RedemptionCode *string `json:"redemption_code,omitempty"`
	SendEmailAt    *string `json:"send_email_at,omitempty"`
}
