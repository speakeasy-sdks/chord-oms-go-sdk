package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListUsersQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListUsersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListUsersPaginationData struct {
	Count       *int64        `json:"count,omitempty"`
	CurrentPage *int64        `json:"current_page,omitempty"`
	Pages       *int64        `json:"pages,omitempty"`
	PerPage     *int64        `json:"per_page,omitempty"`
	TotalCount  *int64        `json:"total_count,omitempty"`
	Users       []shared.User `json:"users,omitempty"`
}

type ListUsers401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListUsersRequest struct {
	QueryParams ListUsersQueryParams
	Security    ListUsersSecurity
}

type ListUsersResponse struct {
	ContentType                       string
	PaginationData                    *ListUsersPaginationData
	StatusCode                        int64
	ListUsers401ApplicationJSONObject *ListUsers401ApplicationJSON
}
