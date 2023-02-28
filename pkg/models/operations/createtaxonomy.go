package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateTaxonomySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateTaxonomyRequest struct {
	Request  shared.TaxonomyInput `request:"mediaType=application/json"`
	Security CreateTaxonomySecurity
}

type CreateTaxonomy422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateTaxonomy401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateTaxonomyResponse struct {
	ContentType                            string
	StatusCode                             int
	CreateTaxonomy401ApplicationJSONObject *CreateTaxonomy401ApplicationJSON
	CreateTaxonomy422ApplicationJSONObject *CreateTaxonomy422ApplicationJSON
	Taxonomy                               *shared.Taxonomy
}
