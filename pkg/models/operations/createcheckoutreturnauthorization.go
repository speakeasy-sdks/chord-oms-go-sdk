package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateCheckoutReturnAuthorizationPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type CreateCheckoutReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateCheckoutReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateCheckoutReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateCheckoutReturnAuthorization422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateCheckoutReturnAuthorizationRequest struct {
	PathParams CreateCheckoutReturnAuthorizationPathParams
	Request    shared.ReturnAuthorizationInput `request:"mediaType=application/json"`
	Security   CreateCheckoutReturnAuthorizationSecurity
}

type CreateCheckoutReturnAuthorizationResponse struct {
	ContentType                                               string
	StatusCode                                                int64
	CreateCheckoutReturnAuthorization401ApplicationJSONObject *CreateCheckoutReturnAuthorization401ApplicationJSON
	CreateCheckoutReturnAuthorization404ApplicationJSONObject *CreateCheckoutReturnAuthorization404ApplicationJSON
	CreateCheckoutReturnAuthorization422ApplicationJSONObject *CreateCheckoutReturnAuthorization422ApplicationJSON
	ReturnAuthorization                                       map[string]interface{}
}
