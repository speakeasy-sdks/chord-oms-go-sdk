package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListOrderPaymentsPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type ListOrderPaymentsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListOrderPaymentsSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type ListOrderPaymentsPaginationData struct {
	Count       *int64           `json:"count,omitempty"`
	CurrentPage *int64           `json:"current_page,omitempty"`
	Pages       *int64           `json:"pages,omitempty"`
	Payments    []shared.Payment `json:"payments,omitempty"`
	PerPage     *int64           `json:"per_page,omitempty"`
	TotalCount  *int64           `json:"total_count,omitempty"`
}

type ListOrderPayments401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListOrderPayments404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListOrderPaymentsRequest struct {
	PathParams  ListOrderPaymentsPathParams
	QueryParams ListOrderPaymentsQueryParams
	Security    ListOrderPaymentsSecurity
}

type ListOrderPaymentsResponse struct {
	ContentType                               string
	PaginationData                            *ListOrderPaymentsPaginationData
	StatusCode                                int64
	ListOrderPayments401ApplicationJSONObject *ListOrderPayments401ApplicationJSON
	ListOrderPayments404ApplicationJSONObject *ListOrderPayments404ApplicationJSON
}
