package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
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
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListReturnAuthorizationRefundsPaginationData struct {
	Count       *int64          `json:"count,omitempty"`
	CurrentPage *int64          `json:"current_page,omitempty"`
	Pages       *int64          `json:"pages,omitempty"`
	PerPage     *int64          `json:"per_page,omitempty"`
	Refunds     []shared.Refund `json:"refunds,omitempty"`
	TotalCount  *int64          `json:"total_count,omitempty"`
}

type ListReturnAuthorizationRefunds401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListReturnAuthorizationRefunds404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListReturnAuthorizationRefundsRequest struct {
	PathParams  ListReturnAuthorizationRefundsPathParams
	QueryParams ListReturnAuthorizationRefundsQueryParams
	Security    ListReturnAuthorizationRefundsSecurity
}

type ListReturnAuthorizationRefundsResponse struct {
	ContentType                                            string
	PaginationData                                         *ListReturnAuthorizationRefundsPaginationData
	StatusCode                                             int64
	ListReturnAuthorizationRefunds401ApplicationJSONObject *ListReturnAuthorizationRefunds401ApplicationJSON
	ListReturnAuthorizationRefunds404ApplicationJSONObject *ListReturnAuthorizationRefunds404ApplicationJSON
}
