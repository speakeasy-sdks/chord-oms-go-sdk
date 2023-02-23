package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type TransferShipmentItemToLocationRequestBody struct {
	OriginalShipmentNumber *string `json:"original_shipment_number,omitempty"`
	Quantity               *int64  `json:"quantity,omitempty"`
	StockLocationID        *int64  `json:"stock_location_id,omitempty"`
	TargetShipmentNumber   *string `json:"target_shipment_number,omitempty"`
	VariantID              *int64  `json:"variant_id,omitempty"`
}

type TransferShipmentItemToLocationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type TransferShipmentItemToLocationRequest struct {
	Request  TransferShipmentItemToLocationRequestBody `request:"mediaType=application/json"`
	Security TransferShipmentItemToLocationSecurity
}

type TransferShipmentItemToLocation422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type TransferShipmentItemToLocation401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type TransferShipmentItemToLocation200ApplicationJSON struct {
	Message              *string `json:"message,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	TargetShipmentNumber *string `json:"target_shipment_number,omitempty"`
}

type TransferShipmentItemToLocationResponse struct {
	ContentType                                            string
	StatusCode                                             int
	TransferShipmentItemToLocation200ApplicationJSONObject *TransferShipmentItemToLocation200ApplicationJSON
	TransferShipmentItemToLocation401ApplicationJSONObject *TransferShipmentItemToLocation401ApplicationJSON
	TransferShipmentItemToLocation422ApplicationJSONObject *TransferShipmentItemToLocation422ApplicationJSON
}
