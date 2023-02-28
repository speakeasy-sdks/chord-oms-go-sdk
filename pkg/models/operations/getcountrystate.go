package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetCountryStatePathParams struct {
	CountryID string `pathParam:"style=simple,explode=false,name=country_id"`
	ID        string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCountryStateRequest struct {
	PathParams GetCountryStatePathParams
}

type GetCountryState404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCountryState401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCountryStateResponse struct {
	ContentType                             string
	StatusCode                              int
	GetCountryState401ApplicationJSONObject *GetCountryState401ApplicationJSON
	GetCountryState404ApplicationJSONObject *GetCountryState404ApplicationJSON
	State                                   *shared.State
}
