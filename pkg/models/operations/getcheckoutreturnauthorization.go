package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetCheckoutReturnAuthorizationPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCheckoutReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetCheckoutReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCheckoutReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCheckoutReturnAuthorizationRequest struct {
	PathParams GetCheckoutReturnAuthorizationPathParams
	Security   GetCheckoutReturnAuthorizationSecurity
}

type GetCheckoutReturnAuthorizationResponse struct {
	ContentType                                            string
	StatusCode                                             int64
	GetCheckoutReturnAuthorization200ApplicationJSONAny    *interface{}
	GetCheckoutReturnAuthorization401ApplicationJSONObject *GetCheckoutReturnAuthorization401ApplicationJSON
	GetCheckoutReturnAuthorization404ApplicationJSONObject *GetCheckoutReturnAuthorization404ApplicationJSON
}
