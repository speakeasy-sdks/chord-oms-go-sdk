package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateTaxonomyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateTaxonomySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateTaxonomyRequest struct {
	PathParams UpdateTaxonomyPathParams
	Request    shared.TaxonomyInput `request:"mediaType=application/json"`
	Security   UpdateTaxonomySecurity
}

type UpdateTaxonomy422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateTaxonomy404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateTaxonomy401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateTaxonomyResponse struct {
	ContentType                            string
	StatusCode                             int
	Taxonomy                               *shared.Taxonomy
	UpdateTaxonomy401ApplicationJSONObject *UpdateTaxonomy401ApplicationJSON
	UpdateTaxonomy404ApplicationJSONObject *UpdateTaxonomy404ApplicationJSON
	UpdateTaxonomy422ApplicationJSONObject *UpdateTaxonomy422ApplicationJSON
}
