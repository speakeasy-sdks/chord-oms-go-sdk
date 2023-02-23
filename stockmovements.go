package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/utils"
	"net/http"
)

type stockMovements struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newStockMovements(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *stockMovements {
	return &stockMovements{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateStockLocationMovement - Create stock location movement
// Creates a stock movement for a stock location.
func (s *stockMovements) CreateStockLocationMovement(ctx context.Context, request operations.CreateStockLocationMovementRequest) (*operations.CreateStockLocationMovementResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/stock_locations/{stock_location_id}/stock_movements", request.PathParams)

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

	res := &operations.CreateStockLocationMovementResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.StockMovement
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.StockMovement = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateStockLocationMovement401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateStockLocationMovement401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateStockLocationMovement404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateStockLocationMovement404ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateStockLocationMovement422ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateStockLocationMovement422ApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetStockLocationMovement - Get stock location movement
// Retrieves a stock location's movement.
func (s *stockMovements) GetStockLocationMovement(ctx context.Context, request operations.GetStockLocationMovementRequest) (*operations.GetStockLocationMovementResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/stock_locations/{stock_location_id}/stock_movements/{id}", request.PathParams)

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

	res := &operations.GetStockLocationMovementResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.StockMovement
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.StockMovement = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetStockLocationMovement401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetStockLocationMovement401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetStockLocationMovement404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetStockLocationMovement404ApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListStockLocationMovements - List stock location movements
// Lists a stock location's movements.
func (s *stockMovements) ListStockLocationMovements(ctx context.Context, request operations.ListStockLocationMovementsRequest) (*operations.ListStockLocationMovementsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/stock_locations/{stock_location_id}/stock_movements", request.PathParams)

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

	res := &operations.ListStockLocationMovementsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListStockLocationMovementsPaginationData
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PaginationData = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListStockLocationMovements401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListStockLocationMovements401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListStockLocationMovements404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListStockLocationMovements404ApplicationJSONObject = out
		}
	}

	return res, nil
}
