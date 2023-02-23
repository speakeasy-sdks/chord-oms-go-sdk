package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type RemoveAddressFromUserAddressBookPathParams struct {
	UserID int64 `pathParam:"style=simple,explode=false,name=user_id"`
}

type RemoveAddressFromUserAddressBookQueryParams struct {
	AddressID int64 `queryParam:"style=form,explode=true,name=address_id"`
}

type RemoveAddressFromUserAddressBookSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type RemoveAddressFromUserAddressBookRequest struct {
	PathParams  RemoveAddressFromUserAddressBookPathParams
	QueryParams RemoveAddressFromUserAddressBookQueryParams
	Security    RemoveAddressFromUserAddressBookSecurity
}

type RemoveAddressFromUserAddressBook500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type RemoveAddressFromUserAddressBook422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type RemoveAddressFromUserAddressBook404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type RemoveAddressFromUserAddressBook401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type RemoveAddressFromUserAddressBookResponse struct {
	ContentType                                              string
	StatusCode                                               int
	RemoveAddressFromUserAddressBook401ApplicationJSONObject *RemoveAddressFromUserAddressBook401ApplicationJSON
	RemoveAddressFromUserAddressBook404ApplicationJSONObject *RemoveAddressFromUserAddressBook404ApplicationJSON
	RemoveAddressFromUserAddressBook422ApplicationJSONObject *RemoveAddressFromUserAddressBook422ApplicationJSON
	RemoveAddressFromUserAddressBook500ApplicationJSONObject *RemoveAddressFromUserAddressBook500ApplicationJSON
}
