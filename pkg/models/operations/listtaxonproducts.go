package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListTaxonProductsQueryParams struct {
	ID      *int64 `queryParam:"style=form,explode=true,name=id"`
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
	Simple  *bool  `queryParam:"style=form,explode=true,name=simple"`
}

type ListTaxonProductsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListTaxonProductsPaginationData struct {
	Count       *int64           `json:"count,omitempty"`
	CurrentPage *int64           `json:"current_page,omitempty"`
	Pages       *int64           `json:"pages,omitempty"`
	PerPage     *int64           `json:"per_page,omitempty"`
	Products    []shared.Product `json:"products,omitempty"`
	TotalCount  *int64           `json:"total_count,omitempty"`
}

type ListTaxonProducts401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListTaxonProductsRequest struct {
	QueryParams ListTaxonProductsQueryParams
	Security    ListTaxonProductsSecurity
}

type ListTaxonProductsResponse struct {
	ContentType                               string
	PaginationData                            *ListTaxonProductsPaginationData
	StatusCode                                int64
	ListTaxonProducts401ApplicationJSONObject *ListTaxonProducts401ApplicationJSON
}
