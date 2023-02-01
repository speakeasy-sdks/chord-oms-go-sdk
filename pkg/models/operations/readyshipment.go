package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ReadyShipmentPathParams struct {
	ShipmentNumber string `pathParam:"style=simple,explode=false,name=shipment_number"`
}

type ReadyShipmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ReadyShipment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ReadyShipment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ReadyShipment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type ReadyShipmentRequest struct {
	PathParams ReadyShipmentPathParams
	Security   ReadyShipmentSecurity
}

type ReadyShipmentResponse struct {
	ContentType                           string
	StatusCode                            int64
	ReadyShipment401ApplicationJSONObject *ReadyShipment401ApplicationJSON
	ReadyShipment404ApplicationJSONObject *ReadyShipment404ApplicationJSON
	ReadyShipment422ApplicationJSONObject *ReadyShipment422ApplicationJSON
	Shipment                              *shared.Shipment
}
