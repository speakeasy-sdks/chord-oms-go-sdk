package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListShipmentEstimatedRatesPathParams struct {
	ShipmentNumber string `pathParam:"style=simple,explode=false,name=shipment_number"`
}

type ListShipmentEstimatedRatesSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type ListShipmentEstimatedRates200ApplicationJSON struct {
	ShippingRates []shared.ShippingRate `json:"shipping_rates,omitempty"`
}

type ListShipmentEstimatedRates401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListShipmentEstimatedRates404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListShipmentEstimatedRatesRequest struct {
	PathParams ListShipmentEstimatedRatesPathParams
	Security   ListShipmentEstimatedRatesSecurity
}

type ListShipmentEstimatedRatesResponse struct {
	ContentType                                        string
	StatusCode                                         int64
	ListShipmentEstimatedRates200ApplicationJSONObject *ListShipmentEstimatedRates200ApplicationJSON
	ListShipmentEstimatedRates401ApplicationJSONObject *ListShipmentEstimatedRates401ApplicationJSON
	ListShipmentEstimatedRates404ApplicationJSONObject *ListShipmentEstimatedRates404ApplicationJSON
}
