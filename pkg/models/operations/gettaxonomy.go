package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetTaxonomyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetTaxonomySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetTaxonomyRequest struct {
	PathParams GetTaxonomyPathParams
	Security   GetTaxonomySecurity
}

type GetTaxonomy404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomy401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomyResponse struct {
	ContentType                         string
	StatusCode                          int
	GetTaxonomy401ApplicationJSONObject *GetTaxonomy401ApplicationJSON
	GetTaxonomy404ApplicationJSONObject *GetTaxonomy404ApplicationJSON
	Taxonomy                            *shared.Taxonomy
}
