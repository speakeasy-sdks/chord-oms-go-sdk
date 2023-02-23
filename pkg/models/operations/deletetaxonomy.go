package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteTaxonomyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteTaxonomySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteTaxonomyRequest struct {
	PathParams DeleteTaxonomyPathParams
	Security   DeleteTaxonomySecurity
}

type DeleteTaxonomy422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteTaxonomy404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteTaxonomy401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteTaxonomyResponse struct {
	ContentType                            string
	StatusCode                             int
	DeleteTaxonomy401ApplicationJSONObject *DeleteTaxonomy401ApplicationJSON
	DeleteTaxonomy404ApplicationJSONObject *DeleteTaxonomy404ApplicationJSON
	DeleteTaxonomy422ApplicationJSONObject *DeleteTaxonomy422ApplicationJSON
}
