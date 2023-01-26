package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListCheckoutReturnAuthorizationPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type ListCheckoutReturnAuthorizationQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListCheckoutReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListCheckoutReturnAuthorizationPaginationData struct {
	Count                *int64                   `json:"count,omitempty"`
	CurrentPage          *int64                   `json:"current_page,omitempty"`
	Pages                *int64                   `json:"pages,omitempty"`
	PerPage              *int64                   `json:"per_page,omitempty"`
	ReturnAuthorizations []map[string]interface{} `json:"return_authorizations,omitempty"`
	TotalCount           *int64                   `json:"total_count,omitempty"`
}

type ListCheckoutReturnAuthorization401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListCheckoutReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListCheckoutReturnAuthorizationRequest struct {
	PathParams  ListCheckoutReturnAuthorizationPathParams
	QueryParams ListCheckoutReturnAuthorizationQueryParams
	Security    ListCheckoutReturnAuthorizationSecurity
}

type ListCheckoutReturnAuthorizationResponse struct {
	ContentType                                             string
	PaginationData                                          *ListCheckoutReturnAuthorizationPaginationData
	StatusCode                                              int64
	ListCheckoutReturnAuthorization401ApplicationJSONObject *ListCheckoutReturnAuthorization401ApplicationJSON
	ListCheckoutReturnAuthorization404ApplicationJSONObject *ListCheckoutReturnAuthorization404ApplicationJSON
}
