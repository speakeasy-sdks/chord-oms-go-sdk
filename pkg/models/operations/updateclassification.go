package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateClassificationRequestBody struct {
	Position  *int64 `json:"position,omitempty"`
	ProductID *int64 `json:"product_id,omitempty"`
	TaxonID   *int64 `json:"taxon_id,omitempty"`
}

type UpdateClassificationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateClassification401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateClassification422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateClassificationRequest struct {
	Request  UpdateClassificationRequestBody `request:"mediaType=application/json"`
	Security UpdateClassificationSecurity
}

type UpdateClassificationResponse struct {
	ContentType                                  string
	StatusCode                                   int64
	UpdateClassification200ApplicationJSONObject map[string]interface{}
	UpdateClassification401ApplicationJSONObject *UpdateClassification401ApplicationJSON
	UpdateClassification422ApplicationJSONObject *UpdateClassification422ApplicationJSON
}
