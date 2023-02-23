package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListStoresQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListStoresSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListStoresRequest struct {
	QueryParams ListStoresQueryParams
	Security    ListStoresSecurity
}

type ListStores401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListStores200ApplicationJSON struct {
	Stores []shared.Store `json:"stores,omitempty"`
}

type ListStoresResponse struct {
	ContentType                        string
	StatusCode                         int
	ListStores200ApplicationJSONObject *ListStores200ApplicationJSON
	ListStores401ApplicationJSONObject *ListStores401ApplicationJSON
}
