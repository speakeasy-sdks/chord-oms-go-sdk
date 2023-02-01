package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateOrderCouponCodePathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type CreateOrderCouponCodeSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOrderCouponCode401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderCouponCode404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderCouponCode422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOrderCouponCodeRequest struct {
	PathParams CreateOrderCouponCodePathParams
	Request    shared.CouponCodeInput `request:"mediaType=application/json"`
	Security   CreateOrderCouponCodeSecurity
}

type CreateOrderCouponCodeResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	CouponCodeHandler                             *shared.CouponCodeHandler
	CreateOrderCouponCode401ApplicationJSONObject *CreateOrderCouponCode401ApplicationJSON
	CreateOrderCouponCode404ApplicationJSONObject *CreateOrderCouponCode404ApplicationJSON
	CreateOrderCouponCode422ApplicationJSONObject *CreateOrderCouponCode422ApplicationJSON
}
