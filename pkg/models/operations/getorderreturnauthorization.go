package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetOrderReturnAuthorizationPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type GetOrderReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetOrderReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderReturnAuthorizationRequest struct {
	PathParams GetOrderReturnAuthorizationPathParams
	Security   GetOrderReturnAuthorizationSecurity
}

type GetOrderReturnAuthorizationResponse struct {
	ContentType                                         string
	StatusCode                                          int64
	GetOrderReturnAuthorization401ApplicationJSONObject *GetOrderReturnAuthorization401ApplicationJSON
	GetOrderReturnAuthorization404ApplicationJSONObject *GetOrderReturnAuthorization404ApplicationJSON
	ReturnAuthorization                                 map[string]interface{}
}
