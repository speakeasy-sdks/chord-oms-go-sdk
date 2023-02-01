package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CancelCheckoutReturnAuthorizationPathParams struct {
	CheckoutID            string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ReturnAuthorizationID string `pathParam:"style=simple,explode=false,name=return_authorization_id"`
}

type CancelCheckoutReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CancelCheckoutReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CancelCheckoutReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CancelCheckoutReturnAuthorization422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CancelCheckoutReturnAuthorizationRequest struct {
	PathParams CancelCheckoutReturnAuthorizationPathParams
	Security   CancelCheckoutReturnAuthorizationSecurity
}

type CancelCheckoutReturnAuthorizationResponse struct {
	ContentType                                               string
	StatusCode                                                int64
	CancelCheckoutReturnAuthorization401ApplicationJSONObject *CancelCheckoutReturnAuthorization401ApplicationJSON
	CancelCheckoutReturnAuthorization404ApplicationJSONObject *CancelCheckoutReturnAuthorization404ApplicationJSON
	CancelCheckoutReturnAuthorization422ApplicationJSONObject *CancelCheckoutReturnAuthorization422ApplicationJSON
	ReturnAuthorization                                       map[string]interface{}
}
