package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type AddShipmentItemPathParams struct {
	ShipmentNumber string `pathParam:"style=simple,explode=false,name=shipment_number"`
}

type AddShipmentItemRequestBody struct {
	Quantity  *int64 `json:"quantity,omitempty"`
	VariantID *int64 `json:"variant_id,omitempty"`
}

type AddShipmentItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type AddShipmentItem401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type AddShipmentItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type AddShipmentItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type AddShipmentItemRequest struct {
	PathParams AddShipmentItemPathParams
	Request    AddShipmentItemRequestBody `request:"mediaType=application/json"`
	Security   AddShipmentItemSecurity
}

type AddShipmentItemResponse struct {
	ContentType                             string
	StatusCode                              int64
	AddShipmentItem401ApplicationJSONObject *AddShipmentItem401ApplicationJSON
	AddShipmentItem404ApplicationJSONObject *AddShipmentItem404ApplicationJSON
	AddShipmentItem422ApplicationJSONObject *AddShipmentItem422ApplicationJSON
	Shipment                                *shared.Shipment
}
