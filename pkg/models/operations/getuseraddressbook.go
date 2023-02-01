package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetUserAddressBookPathParams struct {
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type GetUserAddressBookSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetUserAddressBook401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUserAddressBook404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUserAddressBookRequest struct {
	PathParams GetUserAddressBookPathParams
	Security   GetUserAddressBookSecurity
}

type GetUserAddressBookResponse struct {
	ContentType                                string
	StatusCode                                 int64
	AddressBook                                []shared.AddressBook
	GetUserAddressBook401ApplicationJSONObject *GetUserAddressBook401ApplicationJSON
	GetUserAddressBook404ApplicationJSONObject *GetUserAddressBook404ApplicationJSON
}
