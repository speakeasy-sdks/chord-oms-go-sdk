package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListTaxonsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListTaxonsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListTaxonsPaginationData struct {
	Count       *int64         `json:"count,omitempty"`
	CurrentPage *int64         `json:"current_page,omitempty"`
	Pages       *int64         `json:"pages,omitempty"`
	PerPage     *int64         `json:"per_page,omitempty"`
	Taxons      []shared.Taxon `json:"taxons,omitempty"`
	TotalCount  *int64         `json:"total_count,omitempty"`
}

type ListTaxons401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListTaxonsRequest struct {
	QueryParams ListTaxonsQueryParams
	Security    ListTaxonsSecurity
}

type ListTaxonsResponse struct {
	ContentType                        string
	PaginationData                     *ListTaxonsPaginationData
	StatusCode                         int64
	ListTaxons401ApplicationJSONObject *ListTaxons401ApplicationJSON
}
