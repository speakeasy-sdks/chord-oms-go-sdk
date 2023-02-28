package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListCountriesQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListCountriesRequest struct {
	QueryParams ListCountriesQueryParams
}

type ListCountries401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListCountriesPaginationData struct {
	Count       *int64           `json:"count,omitempty"`
	Countries   []shared.Country `json:"countries,omitempty"`
	CurrentPage *int64           `json:"current_page,omitempty"`
	Pages       *int64           `json:"pages,omitempty"`
	PerPage     *int64           `json:"per_page,omitempty"`
	TotalCount  *int64           `json:"total_count,omitempty"`
}

type ListCountriesResponse struct {
	ContentType                           string
	PaginationData                        *ListCountriesPaginationData
	StatusCode                            int
	ListCountries401ApplicationJSONObject *ListCountries401ApplicationJSON
}
