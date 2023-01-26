package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListCheckoutPaymentsPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type ListCheckoutPaymentsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListCheckoutPaymentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListCheckoutPaymentsPaginationData struct {
	Count       *int64           `json:"count,omitempty"`
	CurrentPage *int64           `json:"current_page,omitempty"`
	Pages       *int64           `json:"pages,omitempty"`
	Payments    []shared.Payment `json:"payments,omitempty"`
	PerPage     *int64           `json:"per_page,omitempty"`
	TotalCount  *int64           `json:"total_count,omitempty"`
}

type ListCheckoutPayments401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListCheckoutPayments404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListCheckoutPaymentsRequest struct {
	PathParams  ListCheckoutPaymentsPathParams
	QueryParams ListCheckoutPaymentsQueryParams
	Security    ListCheckoutPaymentsSecurity
}

type ListCheckoutPaymentsResponse struct {
	ContentType                                  string
	PaginationData                               *ListCheckoutPaymentsPaginationData
	StatusCode                                   int64
	ListCheckoutPayments401ApplicationJSONObject *ListCheckoutPayments401ApplicationJSON
	ListCheckoutPayments404ApplicationJSONObject *ListCheckoutPayments404ApplicationJSON
}
