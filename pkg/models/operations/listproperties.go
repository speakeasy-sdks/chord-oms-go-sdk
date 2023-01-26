package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListPropertiesQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListPropertiesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListPropertiesPaginationData struct {
	Count       *int64            `json:"count,omitempty"`
	CurrentPage *int64            `json:"current_page,omitempty"`
	Pages       *int64            `json:"pages,omitempty"`
	PerPage     *int64            `json:"per_page,omitempty"`
	Properties  []shared.Property `json:"properties,omitempty"`
	TotalCount  *int64            `json:"total_count,omitempty"`
}

type ListProperties401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListPropertiesRequest struct {
	QueryParams ListPropertiesQueryParams
	Security    ListPropertiesSecurity
}

type ListPropertiesResponse struct {
	ContentType                            string
	PaginationData                         *ListPropertiesPaginationData
	StatusCode                             int64
	ListProperties401ApplicationJSONObject *ListProperties401ApplicationJSON
}
