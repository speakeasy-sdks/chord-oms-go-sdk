package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListUserCreditCardsPathParams struct {
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type ListUserCreditCardsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListUserCreditCardsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListUserCreditCardsPaginationData struct {
	Count       *int64              `json:"count,omitempty"`
	CreditCards []shared.CreditCard `json:"credit_cards,omitempty"`
	CurrentPage *int64              `json:"current_page,omitempty"`
	Pages       *int64              `json:"pages,omitempty"`
	PerPage     *int64              `json:"per_page,omitempty"`
	TotalCount  *int64              `json:"total_count,omitempty"`
}

type ListUserCreditCards401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListUserCreditCards404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListUserCreditCardsRequest struct {
	PathParams  ListUserCreditCardsPathParams
	QueryParams ListUserCreditCardsQueryParams
	Security    ListUserCreditCardsSecurity
}

type ListUserCreditCardsResponse struct {
	ContentType                                 string
	PaginationData                              *ListUserCreditCardsPaginationData
	StatusCode                                  int64
	ListUserCreditCards401ApplicationJSONObject *ListUserCreditCards401ApplicationJSON
	ListUserCreditCards404ApplicationJSONObject *ListUserCreditCards404ApplicationJSON
}
