package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListZonesQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListZonesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListZonesPaginationData struct {
	Count       *int64        `json:"count,omitempty"`
	CurrentPage *int64        `json:"current_page,omitempty"`
	Pages       *int64        `json:"pages,omitempty"`
	PerPage     *int64        `json:"per_page,omitempty"`
	TotalCount  *int64        `json:"total_count,omitempty"`
	Zones       []shared.Zone `json:"zones,omitempty"`
}

type ListZones401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListZonesRequest struct {
	QueryParams ListZonesQueryParams
	Security    ListZonesSecurity
}

type ListZonesResponse struct {
	ContentType                       string
	PaginationData                    *ListZonesPaginationData
	StatusCode                        int64
	ListZones401ApplicationJSONObject *ListZones401ApplicationJSON
}
