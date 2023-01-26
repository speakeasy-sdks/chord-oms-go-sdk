package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListTaxonomiesQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListTaxonomiesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListTaxonomiesPaginationData struct {
	Count       *int64            `json:"count,omitempty"`
	CurrentPage *int64            `json:"current_page,omitempty"`
	Pages       *int64            `json:"pages,omitempty"`
	PerPage     *int64            `json:"per_page,omitempty"`
	Taxonomies  []shared.Taxonomy `json:"taxonomies,omitempty"`
	TotalCount  *int64            `json:"total_count,omitempty"`
}

type ListTaxonomies401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListTaxonomiesRequest struct {
	QueryParams ListTaxonomiesQueryParams
	Security    ListTaxonomiesSecurity
}

type ListTaxonomiesResponse struct {
	ContentType                            string
	PaginationData                         *ListTaxonomiesPaginationData
	StatusCode                             int64
	ListTaxonomies401ApplicationJSONObject *ListTaxonomies401ApplicationJSON
}
