package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type RemoveAddressFromUserAddressBookPathParams struct {
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type RemoveAddressFromUserAddressBookQueryParams struct {
	AddressID int64 `queryParam:"style=form,explode=true,name=address_id"`
}

type RemoveAddressFromUserAddressBookSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type RemoveAddressFromUserAddressBook401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type RemoveAddressFromUserAddressBook404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type RemoveAddressFromUserAddressBook422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type RemoveAddressFromUserAddressBookRequest struct {
	PathParams  RemoveAddressFromUserAddressBookPathParams
	QueryParams RemoveAddressFromUserAddressBookQueryParams
	Security    RemoveAddressFromUserAddressBookSecurity
}

type RemoveAddressFromUserAddressBookResponse struct {
	ContentType                                              string
	StatusCode                                               int64
	RemoveAddressFromUserAddressBook401ApplicationJSONObject *RemoveAddressFromUserAddressBook401ApplicationJSON
	RemoveAddressFromUserAddressBook404ApplicationJSONObject *RemoveAddressFromUserAddressBook404ApplicationJSON
	RemoveAddressFromUserAddressBook422ApplicationJSONObject *RemoveAddressFromUserAddressBook422ApplicationJSON
}
