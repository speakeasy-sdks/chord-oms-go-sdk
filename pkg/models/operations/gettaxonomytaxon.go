package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetTaxonomyTaxonPathParams struct {
	ID         string `pathParam:"style=simple,explode=false,name=id"`
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type GetTaxonomyTaxonSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetTaxonomyTaxon401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetTaxonomyTaxon404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetTaxonomyTaxonRequest struct {
	PathParams GetTaxonomyTaxonPathParams
	Security   GetTaxonomyTaxonSecurity
}

type GetTaxonomyTaxonResponse struct {
	ContentType                              string
	StatusCode                               int64
	GetTaxonomyTaxon401ApplicationJSONObject *GetTaxonomyTaxon401ApplicationJSON
	GetTaxonomyTaxon404ApplicationJSONObject *GetTaxonomyTaxon404ApplicationJSON
	Taxon                                    *shared.Taxon
}
