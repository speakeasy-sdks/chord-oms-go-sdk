package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetTaxonJstreePathParams struct {
	TaxonID    string `pathParam:"style=simple,explode=false,name=taxon_id"`
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type GetTaxonJstreeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetTaxonJstree401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetTaxonJstree404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonJstreeRequest struct {
	PathParams GetTaxonJstreePathParams
	Security   GetTaxonJstreeSecurity
}

type GetTaxonJstreeResponse struct {
	ContentType                            string
	StatusCode                             int64
	GetTaxonJstree401ApplicationJSONObject *GetTaxonJstree401ApplicationJSON
	GetTaxonJstree404ApplicationJSONObject *GetTaxonJstree404ApplicationJSON
	Jstrees                                [][]shared.Jstree
}
