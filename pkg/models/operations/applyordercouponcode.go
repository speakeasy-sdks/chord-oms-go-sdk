package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ApplyOrderCouponCodePathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type ApplyOrderCouponCodeSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type ApplyOrderCouponCodeRequest struct {
	PathParams ApplyOrderCouponCodePathParams
	Request    shared.CouponCodeInput `request:"mediaType=application/json"`
	Security   ApplyOrderCouponCodeSecurity
}

type ApplyOrderCouponCode422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type ApplyOrderCouponCode404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ApplyOrderCouponCode401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ApplyOrderCouponCodeResponse struct {
	ContentType                                  string
	StatusCode                                   int
	ApplyOrderCouponCode401ApplicationJSONObject *ApplyOrderCouponCode401ApplicationJSON
	ApplyOrderCouponCode404ApplicationJSONObject *ApplyOrderCouponCode404ApplicationJSON
	ApplyOrderCouponCode422ApplicationJSONObject *ApplyOrderCouponCode422ApplicationJSON
	CouponCodeHandler                            *shared.CouponCodeHandler
}
