package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListTaxonomiesQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListTaxonomiesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListTaxonomiesRequest struct {
	QueryParams ListTaxonomiesQueryParams
	Security    ListTaxonomiesSecurity
}

type ListTaxonomies401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListTaxonomiesPaginationData struct {
	Count       *int64            `json:"count,omitempty"`
	CurrentPage *int64            `json:"current_page,omitempty"`
	Pages       *int64            `json:"pages,omitempty"`
	PerPage     *int64            `json:"per_page,omitempty"`
	Taxonomies  []shared.Taxonomy `json:"taxonomies,omitempty"`
	TotalCount  *int64            `json:"total_count,omitempty"`
}

type ListTaxonomiesResponse struct {
	ContentType                            string
	PaginationData                         *ListTaxonomiesPaginationData
	StatusCode                             int
	ListTaxonomies401ApplicationJSONObject *ListTaxonomies401ApplicationJSON
}
