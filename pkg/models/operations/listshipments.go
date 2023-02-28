package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListShipmentsQueryParams struct {
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Q       *string `queryParam:"style=form,explode=true,name=q"`
}

type ListShipmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListShipmentsRequest struct {
	QueryParams ListShipmentsQueryParams
	Security    ListShipmentsSecurity
}

type ListShipments401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListShipmentsPaginationData struct {
	Count       *int64            `json:"count,omitempty"`
	CurrentPage *int64            `json:"current_page,omitempty"`
	Pages       *int64            `json:"pages,omitempty"`
	PerPage     *int64            `json:"per_page,omitempty"`
	Shipments   []shared.Shipment `json:"shipments,omitempty"`
	TotalCount  *int64            `json:"total_count,omitempty"`
}

type ListShipmentsResponse struct {
	ContentType                           string
	PaginationData                        *ListShipmentsPaginationData
	StatusCode                            int
	ListShipments401ApplicationJSONObject *ListShipments401ApplicationJSON
}
