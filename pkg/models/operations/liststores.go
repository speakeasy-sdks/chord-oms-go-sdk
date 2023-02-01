package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListStoresQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListStoresSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListStores200ApplicationJSON struct {
	Stores []shared.Store `json:"stores,omitempty"`
}

type ListStores401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListStoresRequest struct {
	QueryParams ListStoresQueryParams
	Security    ListStoresSecurity
}

type ListStoresResponse struct {
	ContentType                        string
	StatusCode                         int64
	ListStores200ApplicationJSONObject *ListStores200ApplicationJSON
	ListStores401ApplicationJSONObject *ListStores401ApplicationJSON
}
