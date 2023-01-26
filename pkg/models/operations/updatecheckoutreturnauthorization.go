package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateCheckoutReturnAuthorizationPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateCheckoutReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateCheckoutReturnAuthorization401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateCheckoutReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCheckoutReturnAuthorization422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateCheckoutReturnAuthorizationRequest struct {
	PathParams UpdateCheckoutReturnAuthorizationPathParams
	Request    shared.ReturnAuthorizationInput `request:"mediaType=application/json"`
	Security   UpdateCheckoutReturnAuthorizationSecurity
}

type UpdateCheckoutReturnAuthorizationResponse struct {
	ContentType                                               string
	StatusCode                                                int64
	ReturnAuthorization                                       map[string]interface{}
	UpdateCheckoutReturnAuthorization401ApplicationJSONObject *UpdateCheckoutReturnAuthorization401ApplicationJSON
	UpdateCheckoutReturnAuthorization404ApplicationJSONObject *UpdateCheckoutReturnAuthorization404ApplicationJSON
	UpdateCheckoutReturnAuthorization422ApplicationJSONObject *UpdateCheckoutReturnAuthorization422ApplicationJSON
}
