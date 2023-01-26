package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateShipmentRequestBody struct {
	Quantity        *int64 `json:"quantity,omitempty"`
	StockLocationID *int64 `json:"stock_location_id,omitempty"`
	VariantID       *int64 `json:"variant_id,omitempty"`
}

type CreateShipmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateShipment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateShipment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateShipmentRequest struct {
	Request  CreateShipmentRequestBody `request:"mediaType=application/json"`
	Security CreateShipmentSecurity
}

type CreateShipmentResponse struct {
	ContentType                            string
	StatusCode                             int64
	CreateShipment401ApplicationJSONObject *CreateShipment401ApplicationJSON
	CreateShipment422ApplicationJSONObject *CreateShipment422ApplicationJSON
	Shipment                               *shared.Shipment
}
