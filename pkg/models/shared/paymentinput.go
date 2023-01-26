package shared

type PaymentInputSourceAttributes struct {
	AddressAttributes        *AddressInput `json:"address_attributes,omitempty"`
	CcType                   *string       `json:"cc_type,omitempty"`
	EncryptedData            *string       `json:"encrypted_data,omitempty"`
	ExistingCardID           *int64        `json:"existing_card_id,omitempty"`
	FirstName                *string       `json:"first_name,omitempty"`
	GatewayCustomerProfileID *string       `json:"gateway_customer_profile_id,omitempty"`
	GatewayPaymentProfileID  *string       `json:"gateway_payment_profile_id,omitempty"`
	LastDigits               *string       `json:"last_digits,omitempty"`
	LastName                 *string       `json:"last_name,omitempty"`
	Month                    *int64        `json:"month,omitempty"`
	Name                     *string       `json:"name,omitempty"`
	Number                   *string       `json:"number,omitempty"`
	VerificationValue        *string       `json:"verification_value,omitempty"`
	WalletPaymentSourceID    *int64        `json:"wallet_payment_source_id,omitempty"`
	Year                     *int64        `json:"year,omitempty"`
}

type PaymentInput struct {
	Amount           *string                       `json:"amount,omitempty"`
	PaymentMethod    *string                       `json:"payment_method,omitempty"`
	PaymentMethodID  *int64                        `json:"payment_method_id,omitempty"`
	SourceAttributes *PaymentInputSourceAttributes `json:"source_attributes,omitempty"`
}
