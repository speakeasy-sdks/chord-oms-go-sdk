package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetUserAddressBookPathParams struct {
	UserID int64 `pathParam:"style=simple,explode=false,name=user_id"`
}

type GetUserAddressBookSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetUserAddressBookRequest struct {
	PathParams GetUserAddressBookPathParams
	Security   GetUserAddressBookSecurity
}

type GetUserAddressBook500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type GetUserAddressBook404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUserAddressBook401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUserAddressBookResponse struct {
	ContentType                                string
	StatusCode                                 int
	AddressBook                                []shared.AddressBook
	GetUserAddressBook401ApplicationJSONObject *GetUserAddressBook401ApplicationJSON
	GetUserAddressBook404ApplicationJSONObject *GetUserAddressBook404ApplicationJSON
	GetUserAddressBook500ApplicationJSONObject *GetUserAddressBook500ApplicationJSON
}
