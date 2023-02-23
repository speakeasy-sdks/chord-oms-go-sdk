package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetCheckoutReturnAuthorizationPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCheckoutReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetCheckoutReturnAuthorizationRequest struct {
	PathParams GetCheckoutReturnAuthorizationPathParams
	Security   GetCheckoutReturnAuthorizationSecurity
}

type GetCheckoutReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCheckoutReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCheckoutReturnAuthorizationResponse struct {
	ContentType                                            string
	StatusCode                                             int
	GetCheckoutReturnAuthorization200ApplicationJSONObject map[string]interface{}
	GetCheckoutReturnAuthorization401ApplicationJSONObject *GetCheckoutReturnAuthorization401ApplicationJSON
	GetCheckoutReturnAuthorization404ApplicationJSONObject *GetCheckoutReturnAuthorization404ApplicationJSON
}
