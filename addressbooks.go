package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/utils"
	"net/http"
)

type addressBooks struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newAddressBooks(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *addressBooks {
	return &addressBooks{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetUserAddressBook - Get user address book
// Retrieves a user's address book.
func (s *addressBooks) GetUserAddressBook(ctx context.Context, request operations.GetUserAddressBookRequest) (*operations.GetUserAddressBookResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/users/{user_id}/address_book", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetUserAddressBookResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.AddressBook
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AddressBook = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetUserAddressBook401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetUserAddressBook401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetUserAddressBook404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetUserAddressBook404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetUserAddressBook500ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetUserAddressBook500ApplicationJSONObject = out
		}
	}

	return res, nil
}

// RemoveAddressFromUserAddressBook - Remove address from user address book
// Removes an address from a user's address book.
//
// **Note:** Rather than delete a `Spree::UserAddress` record this action set its `archived` attribute to `true`.
func (s *addressBooks) RemoveAddressFromUserAddressBook(ctx context.Context, request operations.RemoveAddressFromUserAddressBookRequest) (*operations.RemoveAddressFromUserAddressBookResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/users/{user_id}/address_book", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.RemoveAddressFromUserAddressBookResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.RemoveAddressFromUserAddressBook401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.RemoveAddressFromUserAddressBook401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.RemoveAddressFromUserAddressBook404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.RemoveAddressFromUserAddressBook404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.RemoveAddressFromUserAddressBook422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.RemoveAddressFromUserAddressBook422ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.RemoveAddressFromUserAddressBook500ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.RemoveAddressFromUserAddressBook500ApplicationJSONObject = out
		}
	}

	return res, nil
}

// UpdateUserAddressBook - Update user address book
// Updates a user's address book.
//
// **Note:** if the passed `id` matches an existing `address` a new `Spree::Address` record will be created and the matched `address` `archived` on `Spree::UserAddress`. For a similar logic, if the passed `id` matches an existing `address` which is in `archived` state, the `Spree::UserAddress#archived` record will be restored to `false`.
// See `user_address_book.rb` for further information.
func (s *addressBooks) UpdateUserAddressBook(ctx context.Context, request operations.UpdateUserAddressBookRequest) (*operations.UpdateUserAddressBookResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/users/{user_id}/address_book", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateUserAddressBookResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.AddressBook
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AddressBook = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateUserAddressBook401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateUserAddressBook401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateUserAddressBook404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateUserAddressBook404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateUserAddressBook422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateUserAddressBook422ApplicationJSONObject = out
		}
	}

	return res, nil
}
