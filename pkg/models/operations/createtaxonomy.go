package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateTaxonomySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateTaxonomy401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateTaxonomy422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateTaxonomyRequest struct {
	Request  shared.TaxonomyInput `request:"mediaType=application/json"`
	Security CreateTaxonomySecurity
}

type CreateTaxonomyResponse struct {
	ContentType                            string
	StatusCode                             int64
	CreateTaxonomy401ApplicationJSONObject *CreateTaxonomy401ApplicationJSON
	CreateTaxonomy422ApplicationJSONObject *CreateTaxonomy422ApplicationJSON
	Taxonomy                               *shared.Taxonomy
}
