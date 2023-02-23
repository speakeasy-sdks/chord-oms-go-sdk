package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdatePropertyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdatePropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdatePropertyRequest struct {
	PathParams UpdatePropertyPathParams
	Request    *shared.PropertyInput `request:"mediaType=application/json"`
	Security   UpdatePropertySecurity
}

type UpdateProperty422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateProperty401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdatePropertyResponse struct {
	ContentType                            string
	StatusCode                             int
	Property                               *shared.Property
	UpdateProperty401ApplicationJSONObject *UpdateProperty401ApplicationJSON
	UpdateProperty404ApplicationJSONObject *UpdateProperty404ApplicationJSON
	UpdateProperty422ApplicationJSONObject *UpdateProperty422ApplicationJSON
}
