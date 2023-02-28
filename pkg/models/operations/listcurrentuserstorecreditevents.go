package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListCurrentUserStoreCreditEventsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListCurrentUserStoreCreditEventsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListCurrentUserStoreCreditEventsRequest struct {
	QueryParams ListCurrentUserStoreCreditEventsQueryParams
	Security    ListCurrentUserStoreCreditEventsSecurity
}

type ListCurrentUserStoreCreditEvents401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListCurrentUserStoreCreditEventsPaginationData struct {
	Count             *int64                    `json:"count,omitempty"`
	CurrentPage       *int64                    `json:"current_page,omitempty"`
	Pages             *int64                    `json:"pages,omitempty"`
	PerPage           *int64                    `json:"per_page,omitempty"`
	StoreCreditEvents []shared.StoreCreditEvent `json:"store_credit_events,omitempty"`
	TotalCount        *int64                    `json:"total_count,omitempty"`
}

type ListCurrentUserStoreCreditEventsResponse struct {
	ContentType                                              string
	PaginationData                                           *ListCurrentUserStoreCreditEventsPaginationData
	StatusCode                                               int
	ListCurrentUserStoreCreditEvents401ApplicationJSONObject *ListCurrentUserStoreCreditEvents401ApplicationJSON
}
