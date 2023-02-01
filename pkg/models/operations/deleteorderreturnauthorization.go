package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteOrderReturnAuthorizationPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type DeleteOrderReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteOrderReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderReturnAuthorization422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderReturnAuthorizationRequest struct {
	PathParams DeleteOrderReturnAuthorizationPathParams
	Security   DeleteOrderReturnAuthorizationSecurity
}

type DeleteOrderReturnAuthorizationResponse struct {
	ContentType                                            string
	StatusCode                                             int64
	DeleteOrderReturnAuthorization401ApplicationJSONObject *DeleteOrderReturnAuthorization401ApplicationJSON
	DeleteOrderReturnAuthorization404ApplicationJSONObject *DeleteOrderReturnAuthorization404ApplicationJSON
	DeleteOrderReturnAuthorization422ApplicationJSONObject *DeleteOrderReturnAuthorization422ApplicationJSON
	ReturnAuthorization                                    map[string]interface{}
}
