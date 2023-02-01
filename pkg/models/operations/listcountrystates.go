package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListCountryStatesPathParams struct {
	CountryID string `pathParam:"style=simple,explode=false,name=country_id"`
}

type ListCountryStates200ApplicationJSON struct {
	States         []shared.State `json:"states,omitempty"`
	StatesRequired *bool          `json:"states_required,omitempty"`
}

type ListCountryStates401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListCountryStates404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListCountryStatesRequest struct {
	PathParams ListCountryStatesPathParams
}

type ListCountryStatesResponse struct {
	ContentType                               string
	StatusCode                                int64
	ListCountryStates200ApplicationJSONObject *ListCountryStates200ApplicationJSON
	ListCountryStates401ApplicationJSONObject *ListCountryStates401ApplicationJSON
	ListCountryStates404ApplicationJSONObject *ListCountryStates404ApplicationJSON
}
