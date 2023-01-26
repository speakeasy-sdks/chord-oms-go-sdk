package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateTaxonomyTaxonPathParams struct {
	ID         string `pathParam:"style=simple,explode=false,name=id"`
	TaxonomyID string `pathParam:"style=simple,explode=false,name=taxonomy_id"`
}

type UpdateTaxonomyTaxonSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateTaxonomyTaxon401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateTaxonomyTaxon404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateTaxonomyTaxon422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateTaxonomyTaxonRequest struct {
	PathParams UpdateTaxonomyTaxonPathParams
	Request    shared.TaxonInput `request:"mediaType=application/json"`
	Security   UpdateTaxonomyTaxonSecurity
}

type UpdateTaxonomyTaxonResponse struct {
	ContentType                                 string
	StatusCode                                  int64
	Taxon                                       *shared.Taxon
	UpdateTaxonomyTaxon401ApplicationJSONObject *UpdateTaxonomyTaxon401ApplicationJSON
	UpdateTaxonomyTaxon404ApplicationJSONObject *UpdateTaxonomyTaxon404ApplicationJSON
	UpdateTaxonomyTaxon422ApplicationJSONObject *UpdateTaxonomyTaxon422ApplicationJSON
}
