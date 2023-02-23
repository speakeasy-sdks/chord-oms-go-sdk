package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetTaxonomyJstreePathParams struct {
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type GetTaxonomyJstreeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetTaxonomyJstreeRequest struct {
	PathParams GetTaxonomyJstreePathParams
	Security   GetTaxonomyJstreeSecurity
}

type GetTaxonomyJstree404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomyJstree401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomyJstreeResponse struct {
	ContentType                               string
	StatusCode                                int
	GetTaxonomyJstree401ApplicationJSONObject *GetTaxonomyJstree401ApplicationJSON
	GetTaxonomyJstree404ApplicationJSONObject *GetTaxonomyJstree404ApplicationJSON
	Jstree                                    []shared.Jstree
}
