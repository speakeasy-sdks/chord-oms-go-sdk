package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateCheckoutReturnAuthorizationPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type CreateCheckoutReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateCheckoutReturnAuthorizationRequest struct {
	PathParams CreateCheckoutReturnAuthorizationPathParams
	Request    shared.ReturnAuthorizationInput `request:"mediaType=application/json"`
	Security   CreateCheckoutReturnAuthorizationSecurity
}

type CreateCheckoutReturnAuthorization422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateCheckoutReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateCheckoutReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateCheckoutReturnAuthorizationResponse struct {
	ContentType                                               string
	StatusCode                                                int
	CreateCheckoutReturnAuthorization401ApplicationJSONObject *CreateCheckoutReturnAuthorization401ApplicationJSON
	CreateCheckoutReturnAuthorization404ApplicationJSONObject *CreateCheckoutReturnAuthorization404ApplicationJSON
	CreateCheckoutReturnAuthorization422ApplicationJSONObject *CreateCheckoutReturnAuthorization422ApplicationJSON
	ReturnAuthorization                                       map[string]interface{}
}
