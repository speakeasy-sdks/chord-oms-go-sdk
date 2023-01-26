package shared

type PaymentSource struct {
	CcType                   *string `json:"cc_type,omitempty"`
	GatewayCustomerProfileID *string `json:"gateway_customer_profile_id,omitempty"`
	GatewayPaymentProfileID  *string `json:"gateway_payment_profile_id,omitempty"`
	ID                       *int64  `json:"id,omitempty"`
	LastDigits               *string `json:"last_digits,omitempty"`
	Month                    *string `json:"month,omitempty"`
	Name                     *string `json:"name,omitempty"`
	Year                     *string `json:"year,omitempty"`
}
