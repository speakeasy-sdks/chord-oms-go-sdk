package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteTaxonomyTaxonPathParams struct {
	ID         string `pathParam:"style=simple,explode=false,name=id"`
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type DeleteTaxonomyTaxonSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteTaxonomyTaxon401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteTaxonomyTaxon404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteTaxonomyTaxon422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteTaxonomyTaxonRequest struct {
	PathParams DeleteTaxonomyTaxonPathParams
	Security   DeleteTaxonomyTaxonSecurity
}

type DeleteTaxonomyTaxonResponse struct {
	ContentType                                 string
	StatusCode                                  int64
	DeleteTaxonomyTaxon401ApplicationJSONObject *DeleteTaxonomyTaxon401ApplicationJSON
	DeleteTaxonomyTaxon404ApplicationJSONObject *DeleteTaxonomyTaxon404ApplicationJSON
	DeleteTaxonomyTaxon422ApplicationJSONObject *DeleteTaxonomyTaxon422ApplicationJSON
}
