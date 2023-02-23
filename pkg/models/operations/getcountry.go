package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetCountryPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCountryRequest struct {
	PathParams GetCountryPathParams
}

type GetCountry404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCountry401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCountryResponse struct {
	ContentType                        string
	StatusCode                         int
	Country                            *shared.Country
	GetCountry401ApplicationJSONObject *GetCountry401ApplicationJSON
	GetCountry404ApplicationJSONObject *GetCountry404ApplicationJSON
}
