package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type FinalizeCheckoutPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type FinalizeCheckoutQueryParams struct {
	CheckoutSessionID *string `queryParam:"style=form,explode=true,name=checkout_session_id"`
	CheckoutSource    *string `queryParam:"style=form,explode=true,name=checkout_source"`
}

type FinalizeCheckoutSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type FinalizeCheckout401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type FinalizeCheckout404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type FinalizeCheckout422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type FinalizeCheckoutRequest struct {
	PathParams  FinalizeCheckoutPathParams
	QueryParams FinalizeCheckoutQueryParams
	Security    FinalizeCheckoutSecurity
}

type FinalizeCheckoutResponse struct {
	ContentType                              string
	StatusCode                               int64
	FinalizeCheckout401ApplicationJSONObject *FinalizeCheckout401ApplicationJSON
	FinalizeCheckout404ApplicationJSONObject *FinalizeCheckout404ApplicationJSON
	FinalizeCheckout422ApplicationJSONObject *FinalizeCheckout422ApplicationJSON
	OrderBig                                 *shared.OrderBig
}
