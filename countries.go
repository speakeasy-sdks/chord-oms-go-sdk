package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/utils"
	"net/http"
	"strings"
)

type countries struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newCountries(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *countries {
	return &countries{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetCountry - Get country
// Retrieves a country.
func (s *countries) GetCountry(ctx context.Context, request operations.GetCountryRequest) (*operations.GetCountryResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/countries/{id}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.defaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetCountryResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Country
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Country = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetCountry401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetCountry401ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetCountry404ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetCountry404ApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListCountries - List countries
// Lists all available countries.
func (s *countries) ListCountries(ctx context.Context, request operations.ListCountriesRequest) (*operations.ListCountriesResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/countries"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.defaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListCountriesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListCountriesPaginationData
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PaginationData = out
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListCountries401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListCountries401ApplicationJSONObject = out
		}
	}

	return res, nil
}
