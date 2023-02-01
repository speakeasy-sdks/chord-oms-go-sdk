package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteOrderLineItemPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type DeleteOrderLineItemSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteOrderLineItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderLineItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderLineItem422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOrderLineItemRequest struct {
	PathParams DeleteOrderLineItemPathParams
	Security   DeleteOrderLineItemSecurity
}

type DeleteOrderLineItemResponse struct {
	ContentType                                 string
	StatusCode                                  int64
	DeleteOrderLineItem401ApplicationJSONObject *DeleteOrderLineItem401ApplicationJSON
	DeleteOrderLineItem404ApplicationJSONObject *DeleteOrderLineItem404ApplicationJSON
	DeleteOrderLineItem422ApplicationJSONObject *DeleteOrderLineItem422ApplicationJSON
}
