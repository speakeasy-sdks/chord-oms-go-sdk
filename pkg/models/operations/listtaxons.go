package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListTaxonsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListTaxonsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListTaxonsRequest struct {
	QueryParams ListTaxonsQueryParams
	Security    ListTaxonsSecurity
}

type ListTaxons401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListTaxonsPaginationData struct {
	Count       *int64         `json:"count,omitempty"`
	CurrentPage *int64         `json:"current_page,omitempty"`
	Pages       *int64         `json:"pages,omitempty"`
	PerPage     *int64         `json:"per_page,omitempty"`
	Taxons      []shared.Taxon `json:"taxons,omitempty"`
	TotalCount  *int64         `json:"total_count,omitempty"`
}

type ListTaxonsResponse struct {
	ContentType                        string
	PaginationData                     *ListTaxonsPaginationData
	StatusCode                         int
	ListTaxons401ApplicationJSONObject *ListTaxons401ApplicationJSON
}
