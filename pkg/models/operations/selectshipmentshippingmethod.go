package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type SelectShipmentShippingMethodPathParams struct {
	ShipmentNumber string `pathParam:"style=simple,explode=false,name=shipment_number"`
}

type SelectShipmentShippingMethodRequestBody struct {
	ShippingMethodID *int64 `json:"shipping_method_id,omitempty"`
}

type SelectShipmentShippingMethodSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type SelectShipmentShippingMethodRequest struct {
	PathParams SelectShipmentShippingMethodPathParams
	Request    SelectShipmentShippingMethodRequestBody `request:"mediaType=application/json"`
	Security   SelectShipmentShippingMethodSecurity
}

type SelectShipmentShippingMethod404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type SelectShipmentShippingMethod401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type SelectShipmentShippingMethodResponse struct {
	ContentType                                          string
	StatusCode                                           int
	SelectShipmentShippingMethod401ApplicationJSONObject *SelectShipmentShippingMethod401ApplicationJSON
	SelectShipmentShippingMethod404ApplicationJSONObject *SelectShipmentShippingMethod404ApplicationJSON
	Shipment                                             *shared.Shipment
}
