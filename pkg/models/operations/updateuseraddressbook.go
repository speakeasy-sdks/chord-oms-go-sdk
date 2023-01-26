package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateUserAddressBookPathParams struct {
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type UpdateUserAddressBookSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateUserAddressBook401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateUserAddressBook404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateUserAddressBook422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateUserAddressBookRequest struct {
	PathParams UpdateUserAddressBookPathParams
	Request    shared.AddressBookInput `request:"mediaType=application/json"`
	Security   UpdateUserAddressBookSecurity
}

type UpdateUserAddressBookResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	AddressBook                                   []shared.AddressBook
	UpdateUserAddressBook401ApplicationJSONObject *UpdateUserAddressBook401ApplicationJSON
	UpdateUserAddressBook404ApplicationJSONObject *UpdateUserAddressBook404ApplicationJSON
	UpdateUserAddressBook422ApplicationJSONObject *UpdateUserAddressBook422ApplicationJSON
}
