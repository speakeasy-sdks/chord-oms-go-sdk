package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type RemoveShipmentIDPathParams struct {
	ShipmentID string `pathParam:"style=simple,explode=false,name=shipment_id"`
}

type RemoveShipmentIDRequestBody struct {
	Quantity  *int64 `json:"quantity,omitempty"`
	VariantID *int64 `json:"variant_id,omitempty"`
}

type RemoveShipmentIDSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type RemoveShipmentID401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type RemoveShipmentID404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type RemoveShipmentID422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type RemoveShipmentIDRequest struct {
	PathParams RemoveShipmentIDPathParams
	Request    RemoveShipmentIDRequestBody `request:"mediaType=application/json"`
	Security   RemoveShipmentIDSecurity
}

type RemoveShipmentIDResponse struct {
	ContentType                              string
	StatusCode                               int64
	RemoveShipmentID401ApplicationJSONObject *RemoveShipmentID401ApplicationJSON
	RemoveShipmentID404ApplicationJSONObject *RemoveShipmentID404ApplicationJSON
	RemoveShipmentID422ApplicationJSONObject *RemoveShipmentID422ApplicationJSON
	Shipment                                 *shared.Shipment
}
