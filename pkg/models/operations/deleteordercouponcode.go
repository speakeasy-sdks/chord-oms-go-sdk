package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteOrderCouponCodePathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type DeleteOrderCouponCodeSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteOrderCouponCodeRequest struct {
	PathParams DeleteOrderCouponCodePathParams
	Security   DeleteOrderCouponCodeSecurity
}

type DeleteOrderCouponCode422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderCouponCode404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderCouponCode401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderCouponCodeResponse struct {
	ContentType                                   string
	StatusCode                                    int
	DeleteOrderCouponCode401ApplicationJSONObject *DeleteOrderCouponCode401ApplicationJSON
	DeleteOrderCouponCode404ApplicationJSONObject *DeleteOrderCouponCode404ApplicationJSON
	DeleteOrderCouponCode422ApplicationJSONObject *DeleteOrderCouponCode422ApplicationJSON
}
