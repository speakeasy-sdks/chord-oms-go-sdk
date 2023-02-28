package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateTaxonomyTaxonPathParams struct {
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type CreateTaxonomyTaxonSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateTaxonomyTaxonRequest struct {
	PathParams CreateTaxonomyTaxonPathParams
	Request    shared.TaxonInput `request:"mediaType=application/json"`
	Security   CreateTaxonomyTaxonSecurity
}

type CreateTaxonomyTaxon422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateTaxonomyTaxon404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateTaxonomyTaxon401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateTaxonomyTaxonResponse struct {
	ContentType                                 string
	StatusCode                                  int
	CreateTaxonomyTaxon401ApplicationJSONObject *CreateTaxonomyTaxon401ApplicationJSON
	CreateTaxonomyTaxon404ApplicationJSONObject *CreateTaxonomyTaxon404ApplicationJSON
	CreateTaxonomyTaxon422ApplicationJSONObject *CreateTaxonomyTaxon422ApplicationJSON
	Taxon                                       *shared.Taxon
}
