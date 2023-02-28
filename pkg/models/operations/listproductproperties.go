package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListProductPropertiesPathParams struct {
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type ListProductPropertiesQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListProductPropertiesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListProductPropertiesRequest struct {
	PathParams  ListProductPropertiesPathParams
	QueryParams ListProductPropertiesQueryParams
	Security    ListProductPropertiesSecurity
}

type ListProductProperties404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListProductProperties401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListProductPropertiesPaginationData struct {
	Count             *int64                   `json:"count,omitempty"`
	CurrentPage       *int64                   `json:"current_page,omitempty"`
	Pages             *int64                   `json:"pages,omitempty"`
	PerPage           *int64                   `json:"per_page,omitempty"`
	ProductProperties []shared.ProductProperty `json:"product_properties,omitempty"`
	TotalCount        *int64                   `json:"total_count,omitempty"`
}

type ListProductPropertiesResponse struct {
	ContentType                                   string
	PaginationData                                *ListProductPropertiesPaginationData
	StatusCode                                    int
	ListProductProperties401ApplicationJSONObject *ListProductProperties401ApplicationJSON
	ListProductProperties404ApplicationJSONObject *ListProductProperties404ApplicationJSON
}
