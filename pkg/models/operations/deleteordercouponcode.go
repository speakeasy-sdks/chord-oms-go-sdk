package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteOrderCouponCodePathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type DeleteOrderCouponCodeSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteOrderCouponCode401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteOrderCouponCode404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderCouponCode422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderCouponCodeRequest struct {
	PathParams DeleteOrderCouponCodePathParams
	Security   DeleteOrderCouponCodeSecurity
}

type DeleteOrderCouponCodeResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	DeleteOrderCouponCode401ApplicationJSONObject *DeleteOrderCouponCode401ApplicationJSON
	DeleteOrderCouponCode404ApplicationJSONObject *DeleteOrderCouponCode404ApplicationJSON
	DeleteOrderCouponCode422ApplicationJSONObject *DeleteOrderCouponCode422ApplicationJSON
}
