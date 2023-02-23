package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListOrderPaymentsPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type ListOrderPaymentsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListOrderPaymentsSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type ListOrderPaymentsRequest struct {
	PathParams  ListOrderPaymentsPathParams
	QueryParams ListOrderPaymentsQueryParams
	Security    ListOrderPaymentsSecurity
}

type ListOrderPayments404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListOrderPayments401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListOrderPaymentsPaginationData struct {
	Count       *int64           `json:"count,omitempty"`
	CurrentPage *int64           `json:"current_page,omitempty"`
	Pages       *int64           `json:"pages,omitempty"`
	Payments    []shared.Payment `json:"payments,omitempty"`
	PerPage     *int64           `json:"per_page,omitempty"`
	TotalCount  *int64           `json:"total_count,omitempty"`
}

type ListOrderPaymentsResponse struct {
	ContentType                               string
	PaginationData                            *ListOrderPaymentsPaginationData
	StatusCode                                int
	ListOrderPayments401ApplicationJSONObject *ListOrderPayments401ApplicationJSON
	ListOrderPayments404ApplicationJSONObject *ListOrderPayments404ApplicationJSON
}
