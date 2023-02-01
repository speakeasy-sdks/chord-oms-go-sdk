package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateCheckoutPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateCheckoutSquareBraintreeCheckoutInputPaymentsAttributesSourceAttributes struct {
	Nonce       *string `json:"nonce,omitempty"`
	PaymentType *string `json:"payment_type,omitempty"`
}

type UpdateCheckoutSquareBraintreeCheckoutInputPaymentsAttributes struct {
	PaymentMethodType *string                                                                       `json:"payment_method_type,omitempty"`
	SourceAttributes  *UpdateCheckoutSquareBraintreeCheckoutInputPaymentsAttributesSourceAttributes `json:"source_attributes,omitempty"`
}

type UpdateCheckoutSquareBraintreeCheckoutInput struct {
	BillAddressAttributes *shared.AddressInput                                           `json:"bill_address_attributes,omitempty"`
	Channel               *string                                                        `json:"channel,omitempty"`
	CouponCode            *string                                                        `json:"coupon_code,omitempty"`
	Email                 *string                                                        `json:"email,omitempty"`
	Metadata              map[string]interface{}                                         `json:"metadata,omitempty"`
	PaymentsAttributes    []UpdateCheckoutSquareBraintreeCheckoutInputPaymentsAttributes `json:"payments_attributes,omitempty"`
	ReferralIdentifier    *string                                                        `json:"referral_identifier,omitempty"`
	ShipAddressAttributes *shared.AddressInput                                           `json:"ship_address_attributes,omitempty"`
	ShipmentsAttributes   []shared.ShipmentInput                                         `json:"shipments_attributes,omitempty"`
	SpecialInstructions   *string                                                        `json:"special_instructions,omitempty"`
	UseBilling            *bool                                                          `json:"use_billing,omitempty"`
}

type UpdateCheckoutSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateCheckout401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCheckout404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCheckout422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateCheckoutRequest struct {
	PathParams UpdateCheckoutPathParams
	Request    UpdateCheckoutSquareBraintreeCheckoutInput `request:"mediaType=application/json"`
	Security   UpdateCheckoutSecurity
}

type UpdateCheckoutResponse struct {
	ContentType                            string
	StatusCode                             int64
	OrderBig                               *shared.OrderBig
	UpdateCheckout401ApplicationJSONObject *UpdateCheckout401ApplicationJSON
	UpdateCheckout404ApplicationJSONObject *UpdateCheckout404ApplicationJSON
	UpdateCheckout422ApplicationJSONObject *UpdateCheckout422ApplicationJSON
}
