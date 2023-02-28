package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetOrderReturnAuthorizationPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type GetOrderReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetOrderReturnAuthorizationRequest struct {
	PathParams GetOrderReturnAuthorizationPathParams
	Security   GetOrderReturnAuthorizationSecurity
}

type GetOrderReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderReturnAuthorizationResponse struct {
	ContentType                                         string
	StatusCode                                          int
	GetOrderReturnAuthorization401ApplicationJSONObject *GetOrderReturnAuthorization401ApplicationJSON
	GetOrderReturnAuthorization404ApplicationJSONObject *GetOrderReturnAuthorization404ApplicationJSON
	ReturnAuthorization                                 map[string]interface{}
}
