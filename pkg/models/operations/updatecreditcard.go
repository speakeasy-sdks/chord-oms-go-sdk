package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateCreditCardPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateCreditCardSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateCreditCard401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCreditCard404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCreditCard422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateCreditCardRequest struct {
	PathParams UpdateCreditCardPathParams
	Request    shared.CreditCardUpdateInput `request:"mediaType=application/json"`
	Security   UpdateCreditCardSecurity
}

type UpdateCreditCardResponse struct {
	ContentType                              string
	StatusCode                               int64
	CreditCard                               *shared.CreditCard
	UpdateCreditCard401ApplicationJSONObject *UpdateCreditCard401ApplicationJSON
	UpdateCreditCard404ApplicationJSONObject *UpdateCreditCard404ApplicationJSON
	UpdateCreditCard422ApplicationJSONObject *UpdateCreditCard422ApplicationJSON
}
