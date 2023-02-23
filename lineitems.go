package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/utils"
	"net/http"
)

type lineItems struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newLineItems(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *lineItems {
	return &lineItems{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateCheckoutLineItem - Create checkout line item
// Creates a new line item for a checkout.
func (s *lineItems) CreateCheckoutLineItem(ctx context.Context, request operations.CreateCheckoutLineItemRequest) (*operations.CreateCheckoutLineItemResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/line_items", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
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

	res := &operations.CreateCheckoutLineItemResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LineItem
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LineItem = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateCheckoutLineItem401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateCheckoutLineItem401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateCheckoutLineItem404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateCheckoutLineItem404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateCheckoutLineItem422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateCheckoutLineItem422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CreateOrderLineItem - Create order line item
// Creates a line item for an order.
func (s *lineItems) CreateOrderLineItem(ctx context.Context, request operations.CreateOrderLineItemRequest) (*operations.CreateOrderLineItemResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/line_items", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
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

	res := &operations.CreateOrderLineItemResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LineItem
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LineItem = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateOrderLineItem401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateOrderLineItem401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateOrderLineItem404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateOrderLineItem404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateOrderLineItem422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateOrderLineItem422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CrossSellOrderLineItem - Cross-sell order line items
// Cross-sells order line items.
func (s *lineItems) CrossSellOrderLineItem(ctx context.Context, request operations.CrossSellOrderLineItemRequest) (*operations.CrossSellOrderLineItemResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/cross_sell", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
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

	res := &operations.CrossSellOrderLineItemResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LineItem
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LineItem = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CrossSellOrderLineItem401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CrossSellOrderLineItem401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CrossSellOrderLineItem404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CrossSellOrderLineItem404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 405:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CrossSellOrderLineItem405ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CrossSellOrderLineItem405ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CrossSellOrderLineItem422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CrossSellOrderLineItem422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// DeleteCheckoutLineItem - Delete checkout line item
// Deletes a checkout's line item.
func (s *lineItems) DeleteCheckoutLineItem(ctx context.Context, request operations.DeleteCheckoutLineItemRequest) (*operations.DeleteCheckoutLineItemResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/line_items/{id}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
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

	res := &operations.DeleteCheckoutLineItemResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteCheckoutLineItem401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteCheckoutLineItem401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteCheckoutLineItem404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteCheckoutLineItem404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteCheckoutLineItem422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteCheckoutLineItem422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// DeleteOrderLineItem - Delete order line item
// Deletes an order's line item.
func (s *lineItems) DeleteOrderLineItem(ctx context.Context, request operations.DeleteOrderLineItemRequest) (*operations.DeleteOrderLineItemResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/line_items/{id}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
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

	res := &operations.DeleteOrderLineItemResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteOrderLineItem401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteOrderLineItem401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteOrderLineItem404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteOrderLineItem404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteOrderLineItem422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteOrderLineItem422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// UpdateCheckoutLineItem - Update checkout line item
// Updates a checkout's line item.
func (s *lineItems) UpdateCheckoutLineItem(ctx context.Context, request operations.UpdateCheckoutLineItemRequest) (*operations.UpdateCheckoutLineItemResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/line_items/{id}", request.PathParams)

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

	res := &operations.UpdateCheckoutLineItemResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LineItem
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LineItem = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateCheckoutLineItem401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateCheckoutLineItem401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateCheckoutLineItem404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateCheckoutLineItem404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateCheckoutLineItem422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateCheckoutLineItem422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// UpdateOrderLineItem - Update order line item
// Updates an order's line item.
func (s *lineItems) UpdateOrderLineItem(ctx context.Context, request operations.UpdateOrderLineItemRequest) (*operations.UpdateOrderLineItemResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/line_items/{id}", request.PathParams)

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

	res := &operations.UpdateOrderLineItemResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LineItem
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LineItem = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateOrderLineItem401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateOrderLineItem401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateOrderLineItem404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateOrderLineItem404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateOrderLineItem422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateOrderLineItem422ApplicationJSONObject = out
		}
	}

	return res, nil
}
