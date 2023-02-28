package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateUserAddressBookPathParams struct {
	UserID int64 `pathParam:"style=simple,explode=false,name=user_id"`
}

type UpdateUserAddressBookSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateUserAddressBookRequest struct {
	PathParams UpdateUserAddressBookPathParams
	Request    shared.AddressBookInput `request:"mediaType=application/json"`
	Security   UpdateUserAddressBookSecurity
}

type UpdateUserAddressBook422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateUserAddressBook404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateUserAddressBook401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateUserAddressBookResponse struct {
	ContentType                                   string
	StatusCode                                    int
	AddressBook                                   []shared.AddressBook
	UpdateUserAddressBook401ApplicationJSONObject *UpdateUserAddressBook401ApplicationJSON
	UpdateUserAddressBook404ApplicationJSONObject *UpdateUserAddressBook404ApplicationJSON
	UpdateUserAddressBook422ApplicationJSONObject *UpdateUserAddressBook422ApplicationJSON
}
