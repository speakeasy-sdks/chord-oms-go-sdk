package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListShipmentEstimatedRatesPathParams struct {
	ShipmentNumber string `pathParam:"style=simple,explode=false,name=shipment_number"`
}

type ListShipmentEstimatedRatesSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type ListShipmentEstimatedRatesRequest struct {
	PathParams ListShipmentEstimatedRatesPathParams
	Security   ListShipmentEstimatedRatesSecurity
}

type ListShipmentEstimatedRates404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListShipmentEstimatedRates401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListShipmentEstimatedRates200ApplicationJSON struct {
	ShippingRates []shared.ShippingRate `json:"shipping_rates,omitempty"`
}

type ListShipmentEstimatedRatesResponse struct {
	ContentType                                        string
	StatusCode                                         int
	ListShipmentEstimatedRates200ApplicationJSONObject *ListShipmentEstimatedRates200ApplicationJSON
	ListShipmentEstimatedRates401ApplicationJSONObject *ListShipmentEstimatedRates401ApplicationJSON
	ListShipmentEstimatedRates404ApplicationJSONObject *ListShipmentEstimatedRates404ApplicationJSON
}
