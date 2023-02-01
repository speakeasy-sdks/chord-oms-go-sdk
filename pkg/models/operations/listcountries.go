package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListCountriesQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListCountriesPaginationData struct {
	Count       *int64           `json:"count,omitempty"`
	Countries   []shared.Country `json:"countries,omitempty"`
	CurrentPage *int64           `json:"current_page,omitempty"`
	Pages       *int64           `json:"pages,omitempty"`
	PerPage     *int64           `json:"per_page,omitempty"`
	TotalCount  *int64           `json:"total_count,omitempty"`
}

type ListCountries401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListCountriesRequest struct {
	QueryParams ListCountriesQueryParams
}

type ListCountriesResponse struct {
	ContentType                           string
	PaginationData                        *ListCountriesPaginationData
	StatusCode                            int64
	ListCountries401ApplicationJSONObject *ListCountries401ApplicationJSON
}
