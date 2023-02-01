package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListTaxonomyTaxonsPathParams struct {
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type ListTaxonomyTaxonsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListTaxonomyTaxonsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListTaxonomyTaxonsPaginationData struct {
	Count       *int64         `json:"count,omitempty"`
	CurrentPage *int64         `json:"current_page,omitempty"`
	Pages       *int64         `json:"pages,omitempty"`
	PerPage     *int64         `json:"per_page,omitempty"`
	Taxons      []shared.Taxon `json:"taxons,omitempty"`
	TotalCount  *int64         `json:"total_count,omitempty"`
}

type ListTaxonomyTaxons401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListTaxonomyTaxons404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListTaxonomyTaxonsRequest struct {
	PathParams  ListTaxonomyTaxonsPathParams
	QueryParams ListTaxonomyTaxonsQueryParams
	Security    ListTaxonomyTaxonsSecurity
}

type ListTaxonomyTaxonsResponse struct {
	ContentType                                string
	PaginationData                             *ListTaxonomyTaxonsPaginationData
	StatusCode                                 int64
	ListTaxonomyTaxons401ApplicationJSONObject *ListTaxonomyTaxons401ApplicationJSON
	ListTaxonomyTaxons404ApplicationJSONObject *ListTaxonomyTaxons404ApplicationJSON
}
