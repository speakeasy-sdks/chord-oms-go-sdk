package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ShipShipmentPathParams struct {
	ShipmentNumber string `pathParam:"style=simple,explode=false,name=shipment_number"`
}

type ShipShipmentRequestBody struct {
	SendMailer *bool `json:"send_mailer,omitempty"`
}

type ShipShipmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ShipShipment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ShipShipment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ShipShipment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type ShipShipmentRequest struct {
	PathParams ShipShipmentPathParams
	Request    ShipShipmentRequestBody `request:"mediaType=application/json"`
	Security   ShipShipmentSecurity
}

type ShipShipmentResponse struct {
	ContentType                          string
	StatusCode                           int64
	ShipShipment401ApplicationJSONObject *ShipShipment401ApplicationJSON
	ShipShipment404ApplicationJSONObject *ShipShipment404ApplicationJSON
	ShipShipment422ApplicationJSONObject *ShipShipment422ApplicationJSON
	Shipment                             *shared.Shipment
}
