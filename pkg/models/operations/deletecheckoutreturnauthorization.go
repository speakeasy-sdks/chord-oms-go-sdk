package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteCheckoutReturnAuthorizationPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteCheckoutReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteCheckoutReturnAuthorization401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteCheckoutReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteCheckoutReturnAuthorization422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteCheckoutReturnAuthorizationRequest struct {
	PathParams DeleteCheckoutReturnAuthorizationPathParams
	Security   DeleteCheckoutReturnAuthorizationSecurity
}

type DeleteCheckoutReturnAuthorizationResponse struct {
	ContentType                                               string
	StatusCode                                                int64
	DeleteCheckoutReturnAuthorization401ApplicationJSONObject *DeleteCheckoutReturnAuthorization401ApplicationJSON
	DeleteCheckoutReturnAuthorization404ApplicationJSONObject *DeleteCheckoutReturnAuthorization404ApplicationJSON
	DeleteCheckoutReturnAuthorization422ApplicationJSONObject *DeleteCheckoutReturnAuthorization422ApplicationJSON
}
