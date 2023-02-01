package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetCountryPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCountry401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCountry404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCountryRequest struct {
	PathParams GetCountryPathParams
}

type GetCountryResponse struct {
	ContentType                        string
	StatusCode                         int64
	Country                            *shared.Country
	GetCountry401ApplicationJSONObject *GetCountry401ApplicationJSON
	GetCountry404ApplicationJSONObject *GetCountry404ApplicationJSON
}
