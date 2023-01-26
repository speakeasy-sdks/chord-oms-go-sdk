package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListProductVariantsPathParams struct {
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type ListProductVariantsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListProductVariantsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListProductVariantsPaginationData struct {
	Count       *int64           `json:"count,omitempty"`
	CurrentPage *int64           `json:"current_page,omitempty"`
	Pages       *int64           `json:"pages,omitempty"`
	PerPage     *int64           `json:"per_page,omitempty"`
	TotalCount  *int64           `json:"total_count,omitempty"`
	Variants    []shared.Variant `json:"variants,omitempty"`
}

type ListProductVariants401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListProductVariants404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListProductVariantsRequest struct {
	PathParams  ListProductVariantsPathParams
	QueryParams ListProductVariantsQueryParams
	Security    ListProductVariantsSecurity
}

type ListProductVariantsResponse struct {
	ContentType                                 string
	PaginationData                              *ListProductVariantsPaginationData
	StatusCode                                  int64
	ListProductVariants401ApplicationJSONObject *ListProductVariants401ApplicationJSON
	ListProductVariants404ApplicationJSONObject *ListProductVariants404ApplicationJSON
}
