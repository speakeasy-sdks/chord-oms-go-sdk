package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListOptionTypeValuesPathParams struct {
	OptionTypeID string `pathParam:"style=simple,explode=false,name=option_type_id"`
}

type ListOptionTypeValuesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListOptionTypeValues401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListOptionTypeValues404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListOptionTypeValuesRequest struct {
	PathParams ListOptionTypeValuesPathParams
	Security   ListOptionTypeValuesSecurity
}

type ListOptionTypeValuesResponse struct {
	ContentType                                  string
	StatusCode                                   int64
	ListOptionTypeValues401ApplicationJSONObject *ListOptionTypeValues401ApplicationJSON
	ListOptionTypeValues404ApplicationJSONObject *ListOptionTypeValues404ApplicationJSON
	OptionValues                                 []shared.OptionValue
}
