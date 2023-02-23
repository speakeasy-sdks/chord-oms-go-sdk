package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateCreditCardPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateCreditCardSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateCreditCardRequest struct {
	PathParams UpdateCreditCardPathParams
	Request    shared.CreditCardUpdateInput `request:"mediaType=application/json"`
	Security   UpdateCreditCardSecurity
}

type UpdateCreditCard422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateCreditCard404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCreditCard401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCreditCardResponse struct {
	ContentType                              string
	StatusCode                               int
	CreditCard                               *shared.CreditCard
	UpdateCreditCard401ApplicationJSONObject *UpdateCreditCard401ApplicationJSON
	UpdateCreditCard404ApplicationJSONObject *UpdateCreditCard404ApplicationJSON
	UpdateCreditCard422ApplicationJSONObject *UpdateCreditCard422ApplicationJSON
}
