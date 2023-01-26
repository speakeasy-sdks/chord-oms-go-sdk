package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetTaxonomyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetTaxonomySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetTaxonomy401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetTaxonomy404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomyRequest struct {
	PathParams GetTaxonomyPathParams
	Security   GetTaxonomySecurity
}

type GetTaxonomyResponse struct {
	ContentType                         string
	StatusCode                          int64
	GetTaxonomy401ApplicationJSONObject *GetTaxonomy401ApplicationJSON
	GetTaxonomy404ApplicationJSONObject *GetTaxonomy404ApplicationJSON
	Taxonomy                            *shared.Taxonomy
}
