package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListUserShipmentsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListUserShipmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListUserShipmentsPaginationData struct {
	Count       *int64            `json:"count,omitempty"`
	CurrentPage *int64            `json:"current_page,omitempty"`
	Pages       *int64            `json:"pages,omitempty"`
	PerPage     *int64            `json:"per_page,omitempty"`
	Shipments   []shared.Shipment `json:"shipments,omitempty"`
	TotalCount  *int64            `json:"total_count,omitempty"`
}

type ListUserShipments401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListUserShipmentsRequest struct {
	QueryParams ListUserShipmentsQueryParams
	Security    ListUserShipmentsSecurity
}

type ListUserShipmentsResponse struct {
	ContentType                               string
	PaginationData                            *ListUserShipmentsPaginationData
	StatusCode                                int64
	ListUserShipments401ApplicationJSONObject *ListUserShipments401ApplicationJSON
}
