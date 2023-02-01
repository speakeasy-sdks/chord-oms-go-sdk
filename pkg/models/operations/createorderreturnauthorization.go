package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateOrderReturnAuthorizationPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type CreateOrderReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOrderReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderReturnAuthorization422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOrderReturnAuthorizationRequest struct {
	PathParams CreateOrderReturnAuthorizationPathParams
	Request    shared.ReturnAuthorizationInput `request:"mediaType=application/json"`
	Security   CreateOrderReturnAuthorizationSecurity
}

type CreateOrderReturnAuthorizationResponse struct {
	ContentType                                            string
	StatusCode                                             int64
	CreateOrderReturnAuthorization401ApplicationJSONObject *CreateOrderReturnAuthorization401ApplicationJSON
	CreateOrderReturnAuthorization404ApplicationJSONObject *CreateOrderReturnAuthorization404ApplicationJSON
	CreateOrderReturnAuthorization422ApplicationJSONObject *CreateOrderReturnAuthorization422ApplicationJSON
	ReturnAuthorization                                    map[string]interface{}
}
