package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type TransferShipmentItemToShipmentRequestBody struct {
	OriginalShipmentNumber *string `json:"original_shipment_number,omitempty"`
	Quantity               *int64  `json:"quantity,omitempty"`
	TargetShipmentNumber   *string `json:"target_shipment_number,omitempty"`
	VariantID              *int64  `json:"variant_id,omitempty"`
}

type TransferShipmentItemToShipmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type TransferShipmentItemToShipment200ApplicationJSON struct {
	Message              *string `json:"message,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	TargetShipmentNumber *string `json:"target_shipment_number,omitempty"`
}

type TransferShipmentItemToShipment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type TransferShipmentItemToShipment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type TransferShipmentItemToShipmentRequest struct {
	Request  TransferShipmentItemToShipmentRequestBody `request:"mediaType=application/json"`
	Security TransferShipmentItemToShipmentSecurity
}

type TransferShipmentItemToShipmentResponse struct {
	ContentType                                            string
	StatusCode                                             int64
	TransferShipmentItemToShipment200ApplicationJSONObject *TransferShipmentItemToShipment200ApplicationJSON
	TransferShipmentItemToShipment401ApplicationJSONObject *TransferShipmentItemToShipment401ApplicationJSON
	TransferShipmentItemToShipment422ApplicationJSONObject *TransferShipmentItemToShipment422ApplicationJSON
}
