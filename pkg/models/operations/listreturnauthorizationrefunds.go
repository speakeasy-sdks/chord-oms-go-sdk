package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListReturnAuthorizationRefundsPathParams struct {
	ReturnAuthorizationNumber string `pathParam:"style=simple,explode=false,name=return_authorization_number"`
}

type ListReturnAuthorizationRefundsQueryParams struct {
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Q       *string `queryParam:"style=form,explode=true,name=q"`
}

type ListReturnAuthorizationRefundsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListReturnAuthorizationRefundsRequest struct {
	PathParams  ListReturnAuthorizationRefundsPathParams
	QueryParams ListReturnAuthorizationRefundsQueryParams
	Security    ListReturnAuthorizationRefundsSecurity
}

type ListReturnAuthorizationRefunds404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListReturnAuthorizationRefunds401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListReturnAuthorizationRefundsPaginationData struct {
	Count       *int64          `json:"count,omitempty"`
	CurrentPage *int64          `json:"current_page,omitempty"`
	Pages       *int64          `json:"pages,omitempty"`
	PerPage     *int64          `json:"per_page,omitempty"`
	Refunds     []shared.Refund `json:"refunds,omitempty"`
	TotalCount  *int64          `json:"total_count,omitempty"`
}

type ListReturnAuthorizationRefundsResponse struct {
	ContentType                                            string
	PaginationData                                         *ListReturnAuthorizationRefundsPaginationData
	StatusCode                                             int
	ListReturnAuthorizationRefunds401ApplicationJSONObject *ListReturnAuthorizationRefunds401ApplicationJSON
	ListReturnAuthorizationRefunds404ApplicationJSONObject *ListReturnAuthorizationRefunds404ApplicationJSON
}
