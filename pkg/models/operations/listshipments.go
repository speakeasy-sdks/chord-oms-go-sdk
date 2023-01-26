package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListShipmentsQueryParams struct {
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Q       *string `queryParam:"style=form,explode=true,name=q"`
}

type ListShipmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListShipmentsPaginationData struct {
	Count       *int64            `json:"count,omitempty"`
	CurrentPage *int64            `json:"current_page,omitempty"`
	Pages       *int64            `json:"pages,omitempty"`
	PerPage     *int64            `json:"per_page,omitempty"`
	Shipments   []shared.Shipment `json:"shipments,omitempty"`
	TotalCount  *int64            `json:"total_count,omitempty"`
}

type ListShipments401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListShipmentsRequest struct {
	QueryParams ListShipmentsQueryParams
	Security    ListShipmentsSecurity
}

type ListShipmentsResponse struct {
	ContentType                           string
	PaginationData                        *ListShipmentsPaginationData
	StatusCode                            int64
	ListShipments401ApplicationJSONObject *ListShipments401ApplicationJSON
}
