package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/utils"
	"net/http"
)

type payments struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newPayments(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *payments {
	return &payments{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// AuthorizeCheckoutPayment - Authorize checkout payment
// Authorizes a checkout's payment.
func (s *payments) AuthorizeCheckoutPayment(ctx context.Context, request operations.AuthorizeCheckoutPaymentRequest) (*operations.AuthorizeCheckoutPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/payments/{payment_id}/authorize", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.AuthorizeCheckoutPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.AuthorizeCheckoutPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AuthorizeCheckoutPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.AuthorizeCheckoutPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AuthorizeCheckoutPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.AuthorizeCheckoutPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AuthorizeCheckoutPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// AuthorizeOrderPayment - Authorize order payment
// Authorizes an order's payment.
func (s *payments) AuthorizeOrderPayment(ctx context.Context, request operations.AuthorizeOrderPaymentRequest) (*operations.AuthorizeOrderPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/payments/{payment_id}/authorize", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.AuthorizeOrderPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.AuthorizeOrderPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AuthorizeOrderPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.AuthorizeOrderPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AuthorizeOrderPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.AuthorizeOrderPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AuthorizeOrderPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CaptureCheckoutPayment - Capture checkout payment
// Captures a checkout's payment.
func (s *payments) CaptureCheckoutPayment(ctx context.Context, request operations.CaptureCheckoutPaymentRequest) (*operations.CaptureCheckoutPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/payments/{payment_id}/capture", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.CaptureCheckoutPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CaptureCheckoutPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CaptureCheckoutPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CaptureCheckoutPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CaptureCheckoutPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CaptureCheckoutPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CaptureCheckoutPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CaptureOrderPayment - Capture order payment
// Captures an order's payment.
func (s *payments) CaptureOrderPayment(ctx context.Context, request operations.CaptureOrderPaymentRequest) (*operations.CaptureOrderPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/payments/{payment_id}/capture", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.CaptureOrderPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CaptureOrderPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CaptureOrderPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CaptureOrderPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CaptureOrderPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CaptureOrderPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CaptureOrderPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CreateCheckoutPayment - Create checkout payment
// Creates a new payment for a checkout.
func (s *payments) CreateCheckoutPayment(ctx context.Context, request operations.CreateCheckoutPaymentRequest) (*operations.CreateCheckoutPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/payments", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
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

	res := &operations.CreateCheckoutPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateCheckoutPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateCheckoutPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateCheckoutPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateCheckoutPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateCheckoutPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateCheckoutPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CreateOrderPayment - Create order payment
// Creates a payment for an order.
func (s *payments) CreateOrderPayment(ctx context.Context, request operations.CreateOrderPaymentRequest) (*operations.CreateOrderPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/payments", request.PathParams)

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

	res := &operations.CreateOrderPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateOrderPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateOrderPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateOrderPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateOrderPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateOrderPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateOrderPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CreditCheckoutPayment - Credit checkout payment
// Credits a checkout's payment.
func (s *payments) CreditCheckoutPayment(ctx context.Context, request operations.CreditCheckoutPaymentRequest) (*operations.CreditCheckoutPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/payments/{payment_id}/credit", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.CreditCheckoutPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreditCheckoutPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreditCheckoutPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreditCheckoutPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreditCheckoutPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreditCheckoutPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreditCheckoutPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// CreditOrderPayment - Credit order payment
// Credits an order's payment.
func (s *payments) CreditOrderPayment(ctx context.Context, request operations.CreditOrderPaymentRequest) (*operations.CreditOrderPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/payments/{payment_id}/credit", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.CreditOrderPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreditOrderPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreditOrderPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreditOrderPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreditOrderPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreditOrderPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreditOrderPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetCheckoutPayment - Get checkout payment
// Gets a checkout's payment.
func (s *payments) GetCheckoutPayment(ctx context.Context, request operations.GetCheckoutPaymentRequest) (*operations.GetCheckoutPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/payments/{id}", request.PathParams)

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

	res := &operations.GetCheckoutPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetCheckoutPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetCheckoutPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetCheckoutPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetCheckoutPayment404ApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetOrderPayment - Get order payment
// Retrieves an order's payment.
func (s *payments) GetOrderPayment(ctx context.Context, request operations.GetOrderPaymentRequest) (*operations.GetOrderPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/payments/{id}", request.PathParams)

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

	res := &operations.GetOrderPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetOrderPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetOrderPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetOrderPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetOrderPayment404ApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListCheckoutPayments - List checkout payments
// Lists a checkout's payments.
func (s *payments) ListCheckoutPayments(ctx context.Context, request operations.ListCheckoutPaymentsRequest) (*operations.ListCheckoutPaymentsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/payments", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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

	res := &operations.ListCheckoutPaymentsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListCheckoutPaymentsPaginationData
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PaginationData = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListCheckoutPayments401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListCheckoutPayments401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListCheckoutPayments404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListCheckoutPayments404ApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListOrderPayments - List order payments
// Lists an order's payments.
func (s *payments) ListOrderPayments(ctx context.Context, request operations.ListOrderPaymentsRequest) (*operations.ListOrderPaymentsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/payments", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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

	res := &operations.ListOrderPaymentsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListOrderPaymentsPaginationData
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PaginationData = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListOrderPayments401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListOrderPayments401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListOrderPayments404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListOrderPayments404ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PurchaseCheckoutPayment - Purchase checkout payment
// Purchases a checkout's payment.
func (s *payments) PurchaseCheckoutPayment(ctx context.Context, request operations.PurchaseCheckoutPaymentRequest) (*operations.PurchaseCheckoutPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/payments/{payment_id}/purchase", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.PurchaseCheckoutPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PurchaseCheckoutPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PurchaseCheckoutPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PurchaseCheckoutPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PurchaseCheckoutPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PurchaseCheckoutPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PurchaseCheckoutPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PurchaseOrderPayment - Purchase order payment
// Purchases an order's payment.
func (s *payments) PurchaseOrderPayment(ctx context.Context, request operations.PurchaseOrderPaymentRequest) (*operations.PurchaseOrderPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/payments/{payment_id}/purchase", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.PurchaseOrderPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PurchaseOrderPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PurchaseOrderPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PurchaseOrderPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PurchaseOrderPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PurchaseOrderPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PurchaseOrderPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// UpdateCheckoutPayment - Update checkout payment
// Updates a checkout's payment.
func (s *payments) UpdateCheckoutPayment(ctx context.Context, request operations.UpdateCheckoutPaymentRequest) (*operations.UpdateCheckoutPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/payments/{id}", request.PathParams)

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

	res := &operations.UpdateCheckoutPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateCheckoutPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateCheckoutPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateCheckoutPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateCheckoutPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateCheckoutPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateCheckoutPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// UpdateOrderPayment - Update order payment
// Updates an order's payment.
func (s *payments) UpdateOrderPayment(ctx context.Context, request operations.UpdateOrderPaymentRequest) (*operations.UpdateOrderPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/payments/{id}", request.PathParams)

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

	res := &operations.UpdateOrderPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateOrderPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateOrderPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateOrderPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateOrderPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateOrderPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateOrderPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// VoidCheckoutPayment - Void checkout payment
// Voids a checkout's payment.
func (s *payments) VoidCheckoutPayment(ctx context.Context, request operations.VoidCheckoutPaymentRequest) (*operations.VoidCheckoutPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/checkouts/{checkout_id}/payments/{payment_id}/void", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.VoidCheckoutPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.VoidCheckoutPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.VoidCheckoutPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.VoidCheckoutPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.VoidCheckoutPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.VoidCheckoutPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.VoidCheckoutPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// VoidOrderPayment - Void order payment
// Voids an order's payment.
func (s *payments) VoidOrderPayment(ctx context.Context, request operations.VoidOrderPaymentRequest) (*operations.VoidOrderPaymentResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/orders/{order_number}/payments/{payment_id}/void", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
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

	res := &operations.VoidOrderPaymentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Payment
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Payment = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.VoidOrderPayment401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.VoidOrderPayment401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.VoidOrderPayment404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.VoidOrderPayment404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.VoidOrderPayment422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.VoidOrderPayment422ApplicationJSONObject = out
		}
	}

	return res, nil
}
