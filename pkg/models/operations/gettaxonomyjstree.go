package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetTaxonomyJstreePathParams struct {
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type GetTaxonomyJstreeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetTaxonomyJstree401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomyJstree404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomyJstreeRequest struct {
	PathParams GetTaxonomyJstreePathParams
	Security   GetTaxonomyJstreeSecurity
}

type GetTaxonomyJstreeResponse struct {
	ContentType                               string
	StatusCode                                int64
	GetTaxonomyJstree401ApplicationJSONObject *GetTaxonomyJstree401ApplicationJSON
	GetTaxonomyJstree404ApplicationJSONObject *GetTaxonomyJstree404ApplicationJSON
	Jstree                                    []shared.Jstree
}
