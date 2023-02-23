package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateShipmentPathParams struct {
	Number string `pathParam:"style=simple,explode=false,name=number"`
}

type UpdateShipmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateShipmentRequest struct {
	PathParams UpdateShipmentPathParams
	Request    shared.ShipmentInput `request:"mediaType=application/json"`
	Security   UpdateShipmentSecurity
}

type UpdateShipment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateShipment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateShipment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateShipmentResponse struct {
	ContentType                            string
	StatusCode                             int
	Shipment                               *shared.Shipment
	UpdateShipment401ApplicationJSONObject *UpdateShipment401ApplicationJSON
	UpdateShipment404ApplicationJSONObject *UpdateShipment404ApplicationJSON
	UpdateShipment422ApplicationJSONObject *UpdateShipment422ApplicationJSON
}
