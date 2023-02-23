package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetTaxonomyTaxonPathParams struct {
	ID         string `pathParam:"style=simple,explode=false,name=id"`
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type GetTaxonomyTaxonSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetTaxonomyTaxonRequest struct {
	PathParams GetTaxonomyTaxonPathParams
	Security   GetTaxonomyTaxonSecurity
}

type GetTaxonomyTaxon404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomyTaxon401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomyTaxonResponse struct {
	ContentType                              string
	StatusCode                               int
	GetTaxonomyTaxon401ApplicationJSONObject *GetTaxonomyTaxon401ApplicationJSON
	GetTaxonomyTaxon404ApplicationJSONObject *GetTaxonomyTaxon404ApplicationJSON
	Taxon                                    *shared.Taxon
}
