package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetCountryStatePathParams struct {
	CountryID string `pathParam:"style=simple,explode=false,name=country_id"`
	ID        string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCountryState401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetCountryState404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCountryStateRequest struct {
	PathParams GetCountryStatePathParams
}

type GetCountryStateResponse struct {
	ContentType                             string
	StatusCode                              int64
	GetCountryState401ApplicationJSONObject *GetCountryState401ApplicationJSON
	GetCountryState404ApplicationJSONObject *GetCountryState404ApplicationJSON
	State                                   *shared.State
}
