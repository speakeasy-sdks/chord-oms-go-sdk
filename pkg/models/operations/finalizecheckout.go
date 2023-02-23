package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type FinalizeCheckoutPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type FinalizeCheckoutQueryParams struct {
	CheckoutSessionID *string `queryParam:"style=form,explode=true,name=checkout_session_id"`
	CheckoutSource    *string `queryParam:"style=form,explode=true,name=checkout_source"`
}

type FinalizeCheckoutSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type FinalizeCheckoutRequest struct {
	PathParams  FinalizeCheckoutPathParams
	QueryParams FinalizeCheckoutQueryParams
	Security    FinalizeCheckoutSecurity
}

type FinalizeCheckout422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type FinalizeCheckout404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type FinalizeCheckout401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type FinalizeCheckoutResponse struct {
	ContentType                              string
	OrderBig                                 *shared.OrderBig
	StatusCode                               int
	FinalizeCheckout401ApplicationJSONObject *FinalizeCheckout401ApplicationJSON
	FinalizeCheckout404ApplicationJSONObject *FinalizeCheckout404ApplicationJSON
	FinalizeCheckout422ApplicationJSONObject *FinalizeCheckout422ApplicationJSON
}
