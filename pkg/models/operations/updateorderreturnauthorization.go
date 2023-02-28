package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateOrderReturnAuthorizationPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type UpdateOrderReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateOrderReturnAuthorizationRequest struct {
	PathParams UpdateOrderReturnAuthorizationPathParams
	Request    shared.ReturnAuthorizationInput `request:"mediaType=application/json"`
	Security   UpdateOrderReturnAuthorizationSecurity
}

type UpdateOrderReturnAuthorization422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOrderReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderReturnAuthorizationResponse struct {
	ContentType                                            string
	StatusCode                                             int
	ReturnAuthorization                                    map[string]interface{}
	UpdateOrderReturnAuthorization401ApplicationJSONObject *UpdateOrderReturnAuthorization401ApplicationJSON
	UpdateOrderReturnAuthorization404ApplicationJSONObject *UpdateOrderReturnAuthorization404ApplicationJSON
	UpdateOrderReturnAuthorization422ApplicationJSONObject *UpdateOrderReturnAuthorization422ApplicationJSON
}
