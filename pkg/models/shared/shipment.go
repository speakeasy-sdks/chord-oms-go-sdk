package shared

type ShipmentFulfillmentService struct {
	Name *string `json:"name,omitempty"`
}

type Shipment struct {
	Adjustments          []Adjustment                `json:"adjustments,omitempty"`
	Carrier              *string                     `json:"carrier,omitempty"`
	Cost                 *string                     `json:"cost,omitempty"`
	ExternalTrackingURL  *string                     `json:"external_tracking_url,omitempty"`
	FulfillmentService   *ShipmentFulfillmentService `json:"fulfillment_service,omitempty"`
	ID                   *int64                      `json:"id,omitempty"`
	Manifest             [][]ShipmentManifest        `json:"manifest,omitempty"`
	Number               *string                     `json:"number,omitempty"`
	OrderID              *string                     `json:"order_id,omitempty"`
	SelectedShippingRate *ShippingRate               `json:"selected_shipping_rate,omitempty"`
	ShippedAt            *string                     `json:"shipped_at,omitempty"`
	ShippingMethods      []ShippingMethod            `json:"shipping_methods,omitempty"`
	ShippingRates        []ShippingRate              `json:"shipping_rates,omitempty"`
	State                *string                     `json:"state,omitempty"`
	StockLocationName    *string                     `json:"stock_location_name,omitempty"`
	Tracking             *string                     `json:"tracking,omitempty"`
	TrackingURL          *string                     `json:"tracking_url,omitempty"`
}
