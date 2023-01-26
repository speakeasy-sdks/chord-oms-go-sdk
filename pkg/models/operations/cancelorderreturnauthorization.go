package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CancelOrderReturnAuthorizationPathParams struct {
	OrderNumber           string `pathParam:"style=simple,explode=false,name=order_number"`
	ReturnAuthorizationID string `pathParam:"style=simple,explode=false,name=return_authorization_id"`
}

type CancelOrderReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CancelOrderReturnAuthorization401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CancelOrderReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CancelOrderReturnAuthorization422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CancelOrderReturnAuthorizationRequest struct {
	PathParams CancelOrderReturnAuthorizationPathParams
	Security   CancelOrderReturnAuthorizationSecurity
}

type CancelOrderReturnAuthorizationResponse struct {
	ContentType                                            string
	StatusCode                                             int64
	CancelOrderReturnAuthorization401ApplicationJSONObject *CancelOrderReturnAuthorization401ApplicationJSON
	CancelOrderReturnAuthorization404ApplicationJSONObject *CancelOrderReturnAuthorization404ApplicationJSON
	CancelOrderReturnAuthorization422ApplicationJSONObject *CancelOrderReturnAuthorization422ApplicationJSON
	ReturnAuthorization                                    map[string]interface{}
}
