package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListVariantsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListVariantsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListVariantsPaginationData struct {
	Count       *int64          `json:"count,omitempty"`
	CurrentPage *int64          `json:"current_page,omitempty"`
	Pages       *int64          `json:"pages,omitempty"`
	PerPage     *int64          `json:"per_page,omitempty"`
	TotalCount  *int64          `json:"total_count,omitempty"`
	Variants    *shared.Variant `json:"variants,omitempty"`
}

type ListVariants401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListVariantsRequest struct {
	QueryParams ListVariantsQueryParams
	Security    ListVariantsSecurity
}

type ListVariantsResponse struct {
	ContentType                          string
	PaginationData                       *ListVariantsPaginationData
	StatusCode                           int64
	ListVariants401ApplicationJSONObject *ListVariants401ApplicationJSON
}
