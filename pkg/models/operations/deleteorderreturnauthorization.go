package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteOrderReturnAuthorizationPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type DeleteOrderReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteOrderReturnAuthorizationRequest struct {
	PathParams DeleteOrderReturnAuthorizationPathParams
	Security   DeleteOrderReturnAuthorizationSecurity
}

type DeleteOrderReturnAuthorization422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderReturnAuthorizationResponse struct {
	ContentType                                            string
	StatusCode                                             int
	DeleteOrderReturnAuthorization401ApplicationJSONObject *DeleteOrderReturnAuthorization401ApplicationJSON
	DeleteOrderReturnAuthorization404ApplicationJSONObject *DeleteOrderReturnAuthorization404ApplicationJSON
	DeleteOrderReturnAuthorization422ApplicationJSONObject *DeleteOrderReturnAuthorization422ApplicationJSON
	ReturnAuthorization                                    map[string]interface{}
}
