package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetTaxonJstreePathParams struct {
	TaxonID    string `pathParam:"style=simple,explode=false,name=taxon_id"`
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type GetTaxonJstreeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetTaxonJstreeRequest struct {
	PathParams GetTaxonJstreePathParams
	Security   GetTaxonJstreeSecurity
}

type GetTaxonJstree404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonJstree401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonJstreeResponse struct {
	ContentType                            string
	StatusCode                             int
	GetTaxonJstree401ApplicationJSONObject *GetTaxonJstree401ApplicationJSON
	GetTaxonJstree404ApplicationJSONObject *GetTaxonJstree404ApplicationJSON
	Jstrees                                [][]shared.Jstree
}
