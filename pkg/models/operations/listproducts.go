package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListProductsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListProductsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListProductsPaginationData struct {
	Count       *int64           `json:"count,omitempty"`
	CurrentPage *int64           `json:"current_page,omitempty"`
	Pages       *int64           `json:"pages,omitempty"`
	PerPage     *int64           `json:"per_page,omitempty"`
	Products    []shared.Product `json:"products,omitempty"`
	TotalCount  *int64           `json:"total_count,omitempty"`
}

type ListProducts401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListProductsRequest struct {
	QueryParams ListProductsQueryParams
	Security    ListProductsSecurity
}

type ListProductsResponse struct {
	ContentType                          string
	PaginationData                       *ListProductsPaginationData
	StatusCode                           int64
	ListProducts401ApplicationJSONObject *ListProducts401ApplicationJSON
}
