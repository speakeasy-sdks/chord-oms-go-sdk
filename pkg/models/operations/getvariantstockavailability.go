package operations

type GetVariantStockAvailabilityPathParams struct {
	VariantID string `pathParam:"style=simple,explode=false,name=variant_id"`
}

type GetVariantStockAvailabilityQueryParams struct {
	Channel *string `queryParam:"style=form,explode=true,name=channel"`
	Country *string `queryParam:"style=form,explode=true,name=country"`
}

type GetVariantStockAvailability200ApplicationJSON struct {
	ID              *int64  `json:"id,omitempty"`
	InStock         *bool   `json:"in_stock,omitempty"`
	IsBackorderable *bool   `json:"is_backorderable,omitempty"`
	Sku             *string `json:"sku,omitempty"`
}

type GetVariantStockAvailability404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetVariantStockAvailabilityRequest struct {
	PathParams  GetVariantStockAvailabilityPathParams
	QueryParams GetVariantStockAvailabilityQueryParams
}

type GetVariantStockAvailabilityResponse struct {
	ContentType                                         string
	StatusCode                                          int64
	GetVariantStockAvailability200ApplicationJSONObject *GetVariantStockAvailability200ApplicationJSON
	GetVariantStockAvailability404ApplicationJSONObject *GetVariantStockAvailability404ApplicationJSON
}
